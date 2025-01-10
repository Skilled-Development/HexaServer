package play

import (
	"HexaUtils/entities/player"
	"HexaUtils/packets"
	packet_utils "HexaUtils/packets/utils"
)

type SetHeadRotationPacket_1_21 struct {
	PacketID          int
	ClientBoundPacket bool
	ProtocolVersion   int
	State             player.ClientState
	EntityID          int32
	HeadYaw           byte
}

func (p *SetHeadRotationPacket_1_21) GetEntityID() int32 {
	return p.EntityID
}

func (p *SetHeadRotationPacket_1_21) SetEntityID(entityID int32) {
	p.EntityID = entityID
}

func (p *SetHeadRotationPacket_1_21) GetHeadYaw() byte {
	return p.HeadYaw
}

func (p *SetHeadRotationPacket_1_21) SetHeadYaw(headYaw byte) {
	p.HeadYaw = headYaw
}

func (p *SetHeadRotationPacket_1_21) GetProtocolVersion() int {
	return p.ProtocolVersion
}

func (p *SetHeadRotationPacket_1_21) GetPacketID() int {
	return p.PacketID
}

func (p *SetHeadRotationPacket_1_21) GetState() player.ClientState {
	return p.State
}
func (p *SetHeadRotationPacket_1_21) IsClientBound() bool {
	return p.ClientBoundPacket
}

func NewSetHeadRotationPacket_1_21(entityID int32, headYaw byte) SetHeadRotationPacket_1_21 {
	return SetHeadRotationPacket_1_21{
		PacketID:          0x48,
		ClientBoundPacket: true,
		ProtocolVersion:   767, // Or your specific protocol version
		State:             player.Play,
		EntityID:          entityID,
		HeadYaw:           headYaw,
	}
}

func ReadSetHeadRotationPacket_1_21(packet *packet_utils.PacketReader) (*SetHeadRotationPacket_1_21, bool) {
	entityID, err := packet.ReadVarInt()
	if err != nil {
		return &SetHeadRotationPacket_1_21{}, false
	}
	headYaw, err := packet.ReadAngle()
	if err != nil {
		return &SetHeadRotationPacket_1_21{}, false
	}
	return &SetHeadRotationPacket_1_21{
		PacketID:          0x48,
		ClientBoundPacket: true,
		ProtocolVersion:   767,
		State:             player.Play,
		EntityID:          entityID,
		HeadYaw:           headYaw,
	}, true
}

func (p SetHeadRotationPacket_1_21) GetPacket(player player.Player) *packets.Packet {
	//packet := packet_utils.NewPacketWriter()
	packet := player.GetPacketWritter()
	packet.Reset()
	packet.WriteVarInt(int32(p.PacketID))
	packet.WriteVarInt(p.EntityID)
	packet.WriteAngle(p.HeadYaw)

	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"SetHeadRotationPacket_1_21",
		packet.GetPacketBuffer(),
		p.ClientBoundPacket,
		p.State)
	return real_packet
}
