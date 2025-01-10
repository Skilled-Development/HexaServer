package play

import (
	"HexaUtils/entities/player"
	"HexaUtils/packets"
	packet_utils "HexaUtils/packets/utils"

	"github.com/google/uuid"
)

type MetadataType int

const (
	MetadataTypeByte                   MetadataType = 0
	MetadataTypeVarInt                 MetadataType = 1
	MetadataTypeVarLong                MetadataType = 2
	MetadataTypeFloat                  MetadataType = 3
	MetadataTypeString                 MetadataType = 4
	MetadataTypeTextComponent          MetadataType = 5
	MetadataTypeOptionalTextComponent  MetadataType = 6
	MetadataTypeSlot                   MetadataType = 7
	MetadataTypeBoolean                MetadataType = 8
	MetadataTypeRotations              MetadataType = 9
	MetadataTypePosition               MetadataType = 10
	MetadataTypeOptionalPosition       MetadataType = 11
	MetadataTypeDirection              MetadataType = 12
	MetadataTypeOptionalUUID           MetadataType = 13
	MetadataTypeBlockState             MetadataType = 14
	MetadataTypeOptionalBlockState     MetadataType = 15
	MetadataTypeNBT                    MetadataType = 16
	MetadataTypeParticle               MetadataType = 17
	MetadataTypeParticles              MetadataType = 18
	MetadataTypeVillagerData           MetadataType = 19
	MetadataTypeOptionalVarInt         MetadataType = 20
	MetadataTypePose                   MetadataType = 21
	MetadataTypeCatVariant             MetadataType = 22
	MetadataTypeWolfVariant            MetadataType = 23
	MetadataTypeFrogVariant            MetadataType = 24
	MetadataTypeOptionalGlobalPosition MetadataType = 25
	MetadataTypePaintingVariant        MetadataType = 26
	MetadataTypeSnifferState           MetadataType = 27
	MetadataTypeArmadilloState         MetadataType = 28
	MetadataTypeVector3                MetadataType = 29
	MetadataTypeQuaternion             MetadataType = 30
)

type MetadataEntry struct {
	Index int32
	Type  MetadataType
	Value interface{}
}

func ReadMetadataEntry(packet *packet_utils.PacketReader) (*MetadataEntry, bool) {
	index, err := packet.ReadUnsignedByte()
	if err != nil {
		return nil, false
	}

	if index == 0xff {
		return &MetadataEntry{
			Index: int32(index),
		}, true
	}

	metadataType, err := packet.ReadVarInt()
	if err != nil {
		return nil, false
	}

	var value interface{}

	switch MetadataType(metadataType) {
	case MetadataTypeByte:
		val, err := packet.ReadByte()
		if err != nil {
			return nil, false
		}
		value = val
	case MetadataTypeVarInt:
		val, err := packet.ReadVarInt()
		if err != nil {
			return nil, false
		}
		value = val
	case MetadataTypeVarLong:
		val, err := packet.ReadVarLong()
		if err != nil {
			return nil, false
		}
		value = val
	case MetadataTypeFloat:
		val, err := packet.ReadFloat()
		if err != nil {
			return nil, false
		}
		value = val
	case MetadataTypeString:
		val, err := packet.ReadString()
		if err != nil {
			return nil, false
		}
		value = val
	case MetadataTypeTextComponent:
		val, err := packet.ReadString()
		if err != nil {
			return nil, false
		}
		value = val
	case MetadataTypeOptionalTextComponent:
		isPresent, err := packet.ReadBoolean()
		if err != nil {
			return nil, false
		}
		if isPresent {
			textComponent, err := packet.ReadString()
			if err != nil {
				return nil, false
			}
			value = textComponent
		} else {
			value = nil
		}
	/*TODO: Implement Slot

	case MetadataTypeSlot:
			slot, err := packet.ReadSlot()
			if err != nil {
				return nil, false
			}
			value = slot*/
	case MetadataTypeBoolean:
		val, err := packet.ReadBoolean()
		if err != nil {
			return nil, false
		}
		value = val
	case MetadataTypeRotations:
		x, err := packet.ReadFloat()
		if err != nil {
			return nil, false
		}
		y, err := packet.ReadFloat()
		if err != nil {
			return nil, false
		}
		z, err := packet.ReadFloat()
		if err != nil {
			return nil, false
		}

		value = [3]float32{x, y, z}
	/*TODO: Implement Position

	case MetadataTypePosition:
			pos, err := packet.ReadPosition()
			if err != nil {
				return nil, false
			}
			value = pos*/
	/*TODO: Implement OptionalPosition
	case MetadataTypeOptionalPosition:
			isPresent, err := packet.ReadBoolean()
			if err != nil {
				return nil, false
			}
			if isPresent {
				pos, err := packet.ReadPosition()
				if err != nil {
					return nil, false
				}
				value = pos
			} else {
				value = nil
			}*/
	case MetadataTypeDirection:
		val, err := packet.ReadVarInt()
		if err != nil {
			return nil, false
		}
		value = val
	case MetadataTypeOptionalUUID:
		isPresent, err := packet.ReadBoolean()
		if err != nil {
			return nil, false
		}
		if isPresent {
			val, err := packet.ReadUUID()
			if err != nil {
				return nil, false
			}
			value = val
		} else {
			value = nil
		}
	case MetadataTypeBlockState:
		val, err := packet.ReadVarInt()
		if err != nil {
			return nil, false
		}
		value = val
	case MetadataTypeOptionalBlockState:
		val, err := packet.ReadVarInt()
		if err != nil {
			return nil, false
		}
		value = val
	/*TODO: Implement NBT
	case MetadataTypeNBT:
			nbt, err := packet.ReadNBT()
			if err != nil {
				return nil, false
			}
			value = nbt*/
	case MetadataTypeParticle:
		particleType, err := packet.ReadVarInt()
		if err != nil {
			return nil, false
		}

		var particleData interface{}

		// For now this is very basic
		switch int(particleType) {
		case 3:
			//Dust Particle
			color, err := packet.ReadFloat()
			if err != nil {
				return nil, false
			}
			size, err := packet.ReadFloat()
			if err != nil {
				return nil, false
			}

			particleData = [2]float32{color, size}
		}
		value = [2]interface{}{particleType, particleData}
	case MetadataTypeParticles:
		length, err := packet.ReadVarInt()
		if err != nil {
			return nil, false
		}
		particles := make([][2]interface{}, length)
		for i := 0; i < int(length); i++ {
			particleType, err := packet.ReadVarInt()
			if err != nil {
				return nil, false
			}

			var particleData interface{}

			// For now this is very basic
			switch int(particleType) {
			case 3:
				//Dust Particle
				color, err := packet.ReadFloat()
				if err != nil {
					return nil, false
				}
				size, err := packet.ReadFloat()
				if err != nil {
					return nil, false
				}

				particleData = [2]float32{color, size}
			}

			particles[i] = [2]interface{}{particleType, particleData}

		}
		value = particles
	case MetadataTypeVillagerData:
		villagerType, err := packet.ReadVarInt()
		if err != nil {
			return nil, false
		}
		villagerProfession, err := packet.ReadVarInt()
		if err != nil {
			return nil, false
		}
		level, err := packet.ReadVarInt()
		if err != nil {
			return nil, false
		}
		value = [3]int32{villagerType, villagerProfession, level}
	case MetadataTypeOptionalVarInt:
		val, err := packet.ReadVarInt()
		if err != nil {
			return nil, false
		}
		value = val
	case MetadataTypePose:
		val, err := packet.ReadVarInt()
		if err != nil {
			return nil, false
		}
		value = val
	case MetadataTypeCatVariant:
		val, err := packet.ReadVarInt()
		if err != nil {
			return nil, false
		}
		value = val
	case MetadataTypeWolfVariant:
		// TODO Handle Inline Definition
		val, err := packet.ReadString()
		if err != nil {
			return nil, false
		}
		value = val
	case MetadataTypeFrogVariant:
		val, err := packet.ReadVarInt()
		if err != nil {
			return nil, false
		}
		value = val
	/*TODO: Implement OptionalGlobalPosition
	case MetadataTypeOptionalGlobalPosition:
			isPresent, err := packet.ReadBoolean()
			if err != nil {
				return nil, false
			}
			if isPresent {
				identifier, err := packet.ReadString()
				if err != nil {
					return nil, false
				}
				pos, err := packet.ReadPosition()
				if err != nil {
					return nil, false
				}
				value = [2]interface{}{identifier, pos}
			} else {
				value = nil
			}*/

	case MetadataTypePaintingVariant:
		// TODO Handle Inline Definition
		val, err := packet.ReadString()
		if err != nil {
			return nil, false
		}
		value = val
	case MetadataTypeSnifferState:
		val, err := packet.ReadVarInt()
		if err != nil {
			return nil, false
		}
		value = val
	case MetadataTypeArmadilloState:
		val, err := packet.ReadVarInt()
		if err != nil {
			return nil, false
		}
		value = val
	case MetadataTypeVector3:
		x, err := packet.ReadFloat()
		if err != nil {
			return nil, false
		}
		y, err := packet.ReadFloat()
		if err != nil {
			return nil, false
		}
		z, err := packet.ReadFloat()
		if err != nil {
			return nil, false
		}

		value = [3]float32{x, y, z}

	case MetadataTypeQuaternion:
		x, err := packet.ReadFloat()
		if err != nil {
			return nil, false
		}
		y, err := packet.ReadFloat()
		if err != nil {
			return nil, false
		}
		z, err := packet.ReadFloat()
		if err != nil {
			return nil, false
		}
		w, err := packet.ReadFloat()
		if err != nil {
			return nil, false
		}

		value = [4]float32{x, y, z, w}
	default:
		return nil, false
	}

	return &MetadataEntry{
		Index: int32(index),
		Type:  MetadataType(metadataType),
		Value: value,
	}, true
}

func (entry *MetadataEntry) Write(packet *packet_utils.PacketWriter) {
	packet.WriteUnsignedByte(uint8(entry.Index))
	if entry.Index == 0xff {
		return
	}
	packet.WriteVarInt(int32(entry.Type))

	switch entry.Type {
	case MetadataTypeByte:
		packet.WriteByte(byte(entry.Value.(int8)))
	case MetadataTypeVarInt:
		packet.WriteVarInt(entry.Value.(int32))
	case MetadataTypeVarLong:
		packet.WriteVarLong(entry.Value.(int64))
	case MetadataTypeFloat:
		packet.WriteFloat(entry.Value.(float32))
	case MetadataTypeString:
		packet.WriteString(entry.Value.(string))
	case MetadataTypeTextComponent:
		packet.WriteString(entry.Value.(string))
	case MetadataTypeOptionalTextComponent:
		if entry.Value == nil {
			packet.WriteBoolean(false)
		} else {
			packet.WriteBoolean(true)
			packet.WriteString(entry.Value.(string))
		}
	/*case MetadataTypeSlot:
	packet.WriteSlot(entry.Value.(*packets.Slot))*/
	case MetadataTypeBoolean:
		packet.WriteBoolean(entry.Value.(bool))
	case MetadataTypeRotations:
		rotations := entry.Value.([3]float32)
		packet.WriteFloat(rotations[0])
		packet.WriteFloat(rotations[1])
		packet.WriteFloat(rotations[2])
	/*case MetadataTypePosition:
	packet.WritePosition(entry.Value.(packets.Position))*/
	/*case MetadataTypeOptionalPosition:
	if entry.Value == nil {
		packet.WriteBoolean(false)
	} else {
		packet.WriteBoolean(true)
		packet.WritePosition(entry.Value.(packets.Position))
	}*/
	case MetadataTypeDirection:
		packet.WriteVarInt(entry.Value.(int32))
	case MetadataTypeOptionalUUID:
		if entry.Value == nil {
			packet.WriteBoolean(false)
		} else {
			packet.WriteBoolean(true)
			packet.WriteUUID(entry.Value.(uuid.UUID))
		}
	case MetadataTypeBlockState:
		packet.WriteVarInt(entry.Value.(int32))
	case MetadataTypeOptionalBlockState:
		packet.WriteVarInt(entry.Value.(int32))
	/*case MetadataTypeNBT:
	packet.WriteNBT(entry.Value.([]byte))*/
	case MetadataTypeParticle:
		particleData := entry.Value.([2]interface{})
		packet.WriteVarInt(particleData[0].(int32))

		switch int(particleData[0].(int32)) {
		case 3:
			//Dust particle
			data := particleData[1].([2]float32)
			packet.WriteFloat(data[0])
			packet.WriteFloat(data[1])
		}
	case MetadataTypeParticles:
		particleArray := entry.Value.([][2]interface{})
		packet.WriteVarInt(int32(len(particleArray)))
		for _, particleData := range particleArray {
			packet.WriteVarInt(particleData[0].(int32))

			switch int(particleData[0].(int32)) {
			case 3:
				//Dust particle
				data := particleData[1].([2]float32)
				packet.WriteFloat(data[0])
				packet.WriteFloat(data[1])
			}

		}
	case MetadataTypeVillagerData:
		data := entry.Value.([3]int32)
		packet.WriteVarInt(data[0])
		packet.WriteVarInt(data[1])
		packet.WriteVarInt(data[2])
	case MetadataTypeOptionalVarInt:
		packet.WriteVarInt(entry.Value.(int32))
	case MetadataTypePose:
		packet.WriteVarInt(entry.Value.(int32))
	case MetadataTypeCatVariant:
		packet.WriteVarInt(entry.Value.(int32))
	case MetadataTypeWolfVariant:
		// TODO Handle Inline Definition
		packet.WriteString(entry.Value.(string))
	case MetadataTypeFrogVariant:
		packet.WriteVarInt(entry.Value.(int32))
	/*case MetadataTypeOptionalGlobalPosition:
	if entry.Value == nil {
		packet.WriteBoolean(false)
	} else {
		packet.WriteBoolean(true)
		data := entry.Value.([2]interface{})
		packet.WriteString(data[0].(string))
		packet.WritePosition(data[1].(packets.Position))
	}*/
	case MetadataTypePaintingVariant:
		// TODO Handle Inline Definition
		packet.WriteString(entry.Value.(string))
	case MetadataTypeSnifferState:
		packet.WriteVarInt(entry.Value.(int32))
	case MetadataTypeArmadilloState:
		packet.WriteVarInt(entry.Value.(int32))
	case MetadataTypeVector3:
		data := entry.Value.([3]float32)
		packet.WriteFloat(data[0])
		packet.WriteFloat(data[1])
		packet.WriteFloat(data[2])
	case MetadataTypeQuaternion:
		data := entry.Value.([4]float32)
		packet.WriteFloat(data[0])
		packet.WriteFloat(data[1])
		packet.WriteFloat(data[2])
		packet.WriteFloat(data[3])

	}
}

type SetEntityMetadataPacket_1_21 struct {
	PacketID          int
	ClientBoundPacket bool
	ProtocolVersion   int
	State             player.ClientState
	EntityID          int32
	Metadata          []*MetadataEntry
}

func (p *SetEntityMetadataPacket_1_21) GetEntityID() int32 {
	return p.EntityID
}

func (p *SetEntityMetadataPacket_1_21) SetEntityID(entityID int32) {
	p.EntityID = entityID
}

func (p *SetEntityMetadataPacket_1_21) GetMetadata() []*MetadataEntry {
	return p.Metadata
}

func (p *SetEntityMetadataPacket_1_21) SetMetadata(metadata []*MetadataEntry) {
	p.Metadata = metadata
}

func (p *SetEntityMetadataPacket_1_21) GetProtocolVersion() int {
	return p.ProtocolVersion
}

func (p *SetEntityMetadataPacket_1_21) GetPacketID() int {
	return p.PacketID
}

func (p *SetEntityMetadataPacket_1_21) GetState() player.ClientState {
	return p.State
}

func (p *SetEntityMetadataPacket_1_21) IsClientBound() bool {
	return p.ClientBoundPacket
}

func NewSetEntityMetadataPacket_1_21(entityID int32, metadata []*MetadataEntry) SetEntityMetadataPacket_1_21 {
	return SetEntityMetadataPacket_1_21{
		PacketID:          0x58,
		ClientBoundPacket: true,
		ProtocolVersion:   767,
		State:             player.Play,
		EntityID:          entityID,
		Metadata:          metadata,
	}
}

func ReadSetEntityMetadataPacket_1_21(packet *packet_utils.PacketReader) (*SetEntityMetadataPacket_1_21, bool) {
	entityID, err := packet.ReadVarInt()
	if err != nil {
		return &SetEntityMetadataPacket_1_21{}, false
	}
	metadata := make([]*MetadataEntry, 0)
	for {
		entry, ok := ReadMetadataEntry(packet)
		if !ok {
			return &SetEntityMetadataPacket_1_21{}, false
		}
		metadata = append(metadata, entry)
		if entry.Index == 0xff {
			break
		}
	}

	return &SetEntityMetadataPacket_1_21{
		PacketID:          0x58,
		ClientBoundPacket: true,
		ProtocolVersion:   767,
		State:             player.Play,
		EntityID:          entityID,
		Metadata:          metadata,
	}, true
}

func (p SetEntityMetadataPacket_1_21) GetPacket(player player.Player) *packets.Packet {
	//packet := packet_utils.NewPacketWriter()
	packet := player.GetPacketWritter()
	packet.Reset()
	packet.WriteVarInt(int32(p.PacketID))
	packet.WriteVarInt(p.EntityID)

	for _, entry := range p.Metadata {
		entry.Write(packet)
	}

	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"SetEntityMetadataPacket",
		packet.GetPacketBuffer(),
		p.ClientBoundPacket,
		p.State)
	return real_packet
}
