// clientbound/spawnentity.go
package clientbound

import (
	"HexaProtocol_1_21/entities"
	"HexaUtils/entities/player"
	"HexaUtils/packets"

	"github.com/google/uuid"
)

type SpawnEntityPacket_1_21 struct {
	PacketID          int
	ClientBoundPacket bool
	ProtocolVersion   int
	State             player.ClientState
	EntityID          int32
	EntityUUID        uuid.UUID
	Type              entities.EntityType_1_21
	X                 float64
	Y                 float64
	Z                 float64
	Pitch             float32
	Yaw               float32
	HeadYaw           float32
	Data              int32
	VelocityX         int16
	VelocityY         int16
	VelocityZ         int16
}

func (p *SpawnEntityPacket_1_21) GetEntityID() int32 {
	return p.EntityID
}

func (p *SpawnEntityPacket_1_21) SetEntityID(entityID int32) {
	p.EntityID = entityID
}

func (p *SpawnEntityPacket_1_21) GetEntityUUID() uuid.UUID {
	return p.EntityUUID
}

func (p *SpawnEntityPacket_1_21) SetEntityUUID(entityUUID uuid.UUID) {
	p.EntityUUID = entityUUID
}

func (p *SpawnEntityPacket_1_21) GetType() entities.EntityType_1_21 {
	return p.Type
}

func (p *SpawnEntityPacket_1_21) SetType(entityType entities.EntityType_1_21) {
	p.Type = entityType
}

func (p *SpawnEntityPacket_1_21) GetX() float64 {
	return p.X
}

func (p *SpawnEntityPacket_1_21) SetX(x float64) {
	p.X = x
}

func (p *SpawnEntityPacket_1_21) GetY() float64 {
	return p.Y
}

func (p *SpawnEntityPacket_1_21) SetY(y float64) {
	p.Y = y
}

func (p *SpawnEntityPacket_1_21) GetZ() float64 {
	return p.Z
}

func (p *SpawnEntityPacket_1_21) SetZ(z float64) {
	p.Z = z
}

func (p *SpawnEntityPacket_1_21) GetPitch() float32 {
	return p.Pitch
}

func (p *SpawnEntityPacket_1_21) SetPitch(pitch float32) {
	p.Pitch = pitch
}

func (p *SpawnEntityPacket_1_21) GetYaw() float32 {
	return p.Yaw
}

func (p *SpawnEntityPacket_1_21) SetYaw(yaw float32) {
	p.Yaw = yaw
}

func (p *SpawnEntityPacket_1_21) GetHeadYaw() float32 {
	return p.HeadYaw
}

func (p *SpawnEntityPacket_1_21) SetHeadYaw(headYaw float32) {
	p.HeadYaw = headYaw
}

func (p *SpawnEntityPacket_1_21) GetData() int32 {
	return p.Data
}

func (p *SpawnEntityPacket_1_21) SetData(data int32) {
	p.Data = data
}

func (p *SpawnEntityPacket_1_21) GetVelocityX() int16 {
	return p.VelocityX
}

func (p *SpawnEntityPacket_1_21) SetVelocityX(velocityX int16) {
	p.VelocityX = velocityX
}

func (p *SpawnEntityPacket_1_21) GetVelocityY() int16 {
	return p.VelocityY
}

func (p *SpawnEntityPacket_1_21) SetVelocityY(velocityY int16) {
	p.VelocityY = velocityY
}

func (p *SpawnEntityPacket_1_21) GetVelocityZ() int16 {
	return p.VelocityZ
}

func (p *SpawnEntityPacket_1_21) SetVelocityZ(velocityZ int16) {
	p.VelocityZ = velocityZ
}
func (p *SpawnEntityPacket_1_21) GetProtocolVersion() int {
	return p.ProtocolVersion
}

func (p *SpawnEntityPacket_1_21) GetPacketID() int {
	return p.PacketID
}
func (p *SpawnEntityPacket_1_21) GetState() player.ClientState {
	return p.State
}

func (p *SpawnEntityPacket_1_21) IsClientBound() bool {
	return p.ClientBoundPacket
}

func NewSpawnEntityPacket_1_21(entityID int32, entityUUID uuid.UUID, entityType entities.EntityType_1_21, x float64, y float64, z float64, pitch float32, yaw float32, headYaw float32, data int32, velocityX int16, velocityY int16, velocityZ int16) SpawnEntityPacket_1_21 {
	return SpawnEntityPacket_1_21{
		PacketID:          0x01,
		ClientBoundPacket: true,
		ProtocolVersion:   767, // Or your specific protocol version
		State:             player.Play,
		EntityID:          entityID,
		EntityUUID:        entityUUID,
		Type:              entityType,
		X:                 x,
		Y:                 y,
		Z:                 z,
		Pitch:             float32(pitch),
		Yaw:               float32(yaw),
		HeadYaw:           float32(headYaw),
		Data:              data,
		VelocityX:         velocityX,
		VelocityY:         velocityY,
		VelocityZ:         velocityZ,
	}
}

func ReadSpawnEntityPacket_1_21(packet *packets.PacketReader) (*SpawnEntityPacket_1_21, bool) {
	entityID, err := packet.ReadVarInt()
	if err != nil {
		return &SpawnEntityPacket_1_21{}, false
	}
	entityUUID, err := packet.ReadUUID()
	if err != nil {
		return &SpawnEntityPacket_1_21{}, false
	}
	entityType, err := packet.ReadVarInt()
	if err != nil {
		return &SpawnEntityPacket_1_21{}, false
	}
	x, err := packet.ReadDouble()
	if err != nil {
		return &SpawnEntityPacket_1_21{}, false
	}
	y, err := packet.ReadDouble()
	if err != nil {
		return &SpawnEntityPacket_1_21{}, false
	}
	z, err := packet.ReadDouble()
	if err != nil {
		return &SpawnEntityPacket_1_21{}, false
	}
	pitch, err := packet.ReadAngle()
	if err != nil {
		return &SpawnEntityPacket_1_21{}, false
	}
	yaw, err := packet.ReadAngle()
	if err != nil {
		return &SpawnEntityPacket_1_21{}, false
	}
	headYaw, err := packet.ReadAngle()
	if err != nil {
		return &SpawnEntityPacket_1_21{}, false
	}
	data, err := packet.ReadVarInt()
	if err != nil {
		return &SpawnEntityPacket_1_21{}, false
	}
	velocityX, err := packet.ReadShort()
	if err != nil {
		return &SpawnEntityPacket_1_21{}, false
	}
	velocityY, err := packet.ReadShort()
	if err != nil {
		return &SpawnEntityPacket_1_21{}, false
	}
	velocityZ, err := packet.ReadShort()
	if err != nil {
		return &SpawnEntityPacket_1_21{}, false
	}
	return &SpawnEntityPacket_1_21{
		PacketID:          0x01,
		ClientBoundPacket: true,
		ProtocolVersion:   767,
		State:             player.Play,
		EntityID:          entityID,
		EntityUUID:        entityUUID,
		Type:              entities.EntityType_1_21(entityType),
		X:                 x,
		Y:                 y,
		Z:                 z,
		Pitch:             float32(pitch),
		Yaw:               float32(yaw),
		HeadYaw:           float32(headYaw),
		Data:              data,
		VelocityX:         velocityX,
		VelocityY:         velocityY,
		VelocityZ:         velocityZ,
	}, true
}

func (p SpawnEntityPacket_1_21) GetPacket() *packets.Packet {
	packet := packets.NewPacketWriter()
	packet.WriteVarInt(int32(p.PacketID))
	packet.WriteVarInt(p.EntityID)
	packet.WriteUUID(p.EntityUUID)
	packet.WriteVarInt(int32(p.Type))
	packet.WriteDouble(p.X)
	packet.WriteDouble(p.Y)
	packet.WriteDouble(p.Z)
	packet.WriteAngle(byte(p.Pitch))
	packet.WriteAngle(byte(p.Yaw))
	packet.WriteAngle(byte(p.HeadYaw))
	packet.WriteVarInt(p.Data)
	packet.WriteShort(p.VelocityX)
	packet.WriteShort(p.VelocityY)
	packet.WriteShort(p.VelocityZ)

	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"SpawnEntityPacket_1_21",
		packet.GetPacketBuffer(),
		p.ClientBoundPacket,
		p.State)
	return real_packet
}
