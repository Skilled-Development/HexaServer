// clientbound/teleportentity.go
package play

import (
	"HexaUtils/entities/player"
	"HexaUtils/packets"
)

type TeleportEntityPacket_1_21 struct {
	PacketID          int
	ClientBoundPacket bool
	ProtocolVersion   int
	State             player.ClientState
	EntityID          int32
	X                 float64
	Y                 float64
	Z                 float64
	Yaw               byte
	Pitch             byte
	OnGround          bool
}

func (p *TeleportEntityPacket_1_21) GetEntityID() int32 {
	return p.EntityID
}

func (p *TeleportEntityPacket_1_21) SetEntityID(entityID int32) {
	p.EntityID = entityID
}

func (p *TeleportEntityPacket_1_21) GetX() float64 {
	return p.X
}

func (p *TeleportEntityPacket_1_21) SetX(x float64) {
	p.X = x
}

func (p *TeleportEntityPacket_1_21) GetY() float64 {
	return p.Y
}

func (p *TeleportEntityPacket_1_21) SetY(y float64) {
	p.Y = y
}

func (p *TeleportEntityPacket_1_21) GetZ() float64 {
	return p.Z
}

func (p *TeleportEntityPacket_1_21) SetZ(z float64) {
	p.Z = z
}

func (p *TeleportEntityPacket_1_21) GetYaw() byte {
	return p.Yaw
}

func (p *TeleportEntityPacket_1_21) SetYaw(yaw byte) {
	p.Yaw = yaw
}

func (p *TeleportEntityPacket_1_21) GetPitch() byte {
	return p.Pitch
}

func (p *TeleportEntityPacket_1_21) SetPitch(pitch byte) {
	p.Pitch = pitch
}

func (p *TeleportEntityPacket_1_21) GetOnGround() bool {
	return p.OnGround
}

func (p *TeleportEntityPacket_1_21) SetOnGround(onGround bool) {
	p.OnGround = onGround
}

func (p *TeleportEntityPacket_1_21) GetProtocolVersion() int {
	return p.ProtocolVersion
}

func (p *TeleportEntityPacket_1_21) GetPacketID() int {
	return p.PacketID
}

func (p *TeleportEntityPacket_1_21) GetState() player.ClientState {
	return p.State
}

func (p *TeleportEntityPacket_1_21) IsClientBound() bool {
	return p.ClientBoundPacket
}

func NewTeleportEntityPacket_1_21(entityID int32, x float64, y float64, z float64, yaw byte, pitch byte, onGround bool) TeleportEntityPacket_1_21 {
	return TeleportEntityPacket_1_21{
		PacketID:          0x70,
		ClientBoundPacket: true,
		ProtocolVersion:   767,
		State:             player.Play,
		EntityID:          entityID,
		X:                 x,
		Y:                 y,
		Z:                 z,
		Yaw:               yaw,
		Pitch:             pitch,
		OnGround:          onGround,
	}
}

func ReadTeleportEntityPacket_1_21(packet *packets.PacketReader) (TeleportEntityPacket_1_21, bool) {
	entityID, err := packet.ReadVarInt()
	if err != nil {
		return TeleportEntityPacket_1_21{}, false
	}
	x, err := packet.ReadDouble()
	if err != nil {
		return TeleportEntityPacket_1_21{}, false
	}
	y, err := packet.ReadDouble()
	if err != nil {
		return TeleportEntityPacket_1_21{}, false
	}
	z, err := packet.ReadDouble()
	if err != nil {
		return TeleportEntityPacket_1_21{}, false
	}
	yaw, err := packet.ReadAngle()
	if err != nil {
		return TeleportEntityPacket_1_21{}, false
	}
	pitch, err := packet.ReadAngle()
	if err != nil {
		return TeleportEntityPacket_1_21{}, false
	}
	onGround, err := packet.ReadBoolean()
	if err != nil {
		return TeleportEntityPacket_1_21{}, false
	}
	return TeleportEntityPacket_1_21{
		PacketID:          0x70,
		ClientBoundPacket: true,
		ProtocolVersion:   767,
		State:             player.Play,
		EntityID:          entityID,
		X:                 x,
		Y:                 y,
		Z:                 z,
		Yaw:               yaw,
		Pitch:             pitch,
		OnGround:          onGround,
	}, true
}

func (p TeleportEntityPacket_1_21) GetPacket() *packets.Packet {
	packet := packets.NewPacketWriter()
	packet.WriteVarInt(int32(p.PacketID))
	packet.WriteVarInt(p.EntityID)
	packet.WriteDouble(p.X)
	packet.WriteDouble(p.Y)
	packet.WriteDouble(p.Z)
	packet.WriteAngle(p.Yaw)
	packet.WriteAngle(p.Pitch)
	packet.WriteBoolean(p.OnGround)

	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"TeleportEntityPacket_1_21",
		packet.GetPacketBuffer(),
		p.ClientBoundPacket,
		p.State)
	return real_packet
}
