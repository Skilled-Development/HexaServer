package play

import (
	"HexaUtils/entities/player"
	"HexaUtils/packets"
	packet_utils "HexaUtils/packets/utils"
)

type UpdateEntityRotationPacket_1_21 struct {
	PacketID          int
	ClientBoundPacket bool
	ProtocolVersion   int
	State             player.ClientState
	EntityID          int32
	Yaw               byte
	Pitch             byte
	OnGround          bool
}

func (p *UpdateEntityRotationPacket_1_21) GetEntityID() int32 {
	return p.EntityID
}

func (p *UpdateEntityRotationPacket_1_21) SetEntityID(entityID int32) {
	p.EntityID = entityID
}

func (p *UpdateEntityRotationPacket_1_21) GetYaw() byte {
	return p.Yaw
}

func (p *UpdateEntityRotationPacket_1_21) SetYaw(yaw byte) {
	p.Yaw = yaw
}

func (p *UpdateEntityRotationPacket_1_21) GetPitch() byte {
	return p.Pitch
}

func (p *UpdateEntityRotationPacket_1_21) SetPitch(pitch byte) {
	p.Pitch = pitch
}

func (p *UpdateEntityRotationPacket_1_21) GetOnGround() bool {
	return p.OnGround
}

func (p *UpdateEntityRotationPacket_1_21) SetOnGround(onGround bool) {
	p.OnGround = onGround
}

func (p *UpdateEntityRotationPacket_1_21) GetProtocolVersion() int {
	return p.ProtocolVersion
}

func (p *UpdateEntityRotationPacket_1_21) GetPacketID() int {
	return p.PacketID
}

func (p *UpdateEntityRotationPacket_1_21) GetState() player.ClientState {
	return p.State
}

func (p *UpdateEntityRotationPacket_1_21) IsClientBound() bool {
	return p.ClientBoundPacket
}

func NewUpdateEntityRotationPacket_1_21(entityID int32, yaw byte, pitch byte, onGround bool) UpdateEntityRotationPacket_1_21 {
	return UpdateEntityRotationPacket_1_21{
		PacketID:          0x30, // Packet ID for Update Entity Rotation
		ClientBoundPacket: true,
		ProtocolVersion:   767, // Or your specific protocol version
		State:             player.Play,
		EntityID:          entityID,
		Yaw:               yaw,
		Pitch:             pitch,
		OnGround:          onGround,
	}
}

func ReadUpdateEntityRotationPacket_1_21(packet *packet_utils.PacketReader) (*UpdateEntityRotationPacket_1_21, bool) {
	entityID, err := packet.ReadVarInt()
	if err != nil {
		return &UpdateEntityRotationPacket_1_21{}, false
	}
	yaw, err := packet.ReadAngle()
	if err != nil {
		return &UpdateEntityRotationPacket_1_21{}, false
	}
	pitch, err := packet.ReadAngle()
	if err != nil {
		return &UpdateEntityRotationPacket_1_21{}, false
	}
	onGround, err := packet.ReadBoolean()
	if err != nil {
		return &UpdateEntityRotationPacket_1_21{}, false
	}
	return &UpdateEntityRotationPacket_1_21{
		PacketID:          0x30,
		ClientBoundPacket: true,
		ProtocolVersion:   767,
		State:             player.Play,
		EntityID:          entityID,
		Yaw:               yaw,
		Pitch:             pitch,
		OnGround:          onGround,
	}, true
}

func (p UpdateEntityRotationPacket_1_21) GetPacket(player player.Player) *packets.Packet {
	//packet := packet_utils.NewPacketWriter()
	packet := player.GetPacketWritter()
	packet.Reset()
	packet.WriteVarInt(int32(p.PacketID))
	packet.WriteVarInt(p.EntityID)
	packet.WriteAngle(p.Yaw)
	packet.WriteAngle(p.Pitch)
	packet.WriteBoolean(p.OnGround)

	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"UpdateEntityRotationPacket",
		packet.GetPacketBuffer(),
		p.ClientBoundPacket,
		p.State)
	return real_packet
}
