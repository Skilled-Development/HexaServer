// clientbound/updatepositionandrotation.go
package play

import (
	"HexaUtils/entities/player"
	"HexaUtils/packets"
	packet_utils "HexaUtils/packets/utils"
)

type UpdatePositionAndRotationPacket_1_21 struct {
	PacketID          int
	ClientBoundPacket bool
	ProtocolVersion   int
	State             player.ClientState
	EntityID          int32
	DeltaX            int16
	DeltaY            int16
	DeltaZ            int16
	Yaw               byte
	Pitch             byte
	OnGround          bool
}

func (p *UpdatePositionAndRotationPacket_1_21) GetEntityID() int32 {
	return p.EntityID
}

func (p *UpdatePositionAndRotationPacket_1_21) SetEntityID(entityID int32) {
	p.EntityID = entityID
}

func (p *UpdatePositionAndRotationPacket_1_21) GetDeltaX() int16 {
	return p.DeltaX
}

func (p *UpdatePositionAndRotationPacket_1_21) SetDeltaX(deltaX int16) {
	p.DeltaX = deltaX
}

func (p *UpdatePositionAndRotationPacket_1_21) GetDeltaY() int16 {
	return p.DeltaY
}

func (p *UpdatePositionAndRotationPacket_1_21) SetDeltaY(deltaY int16) {
	p.DeltaY = deltaY
}

func (p *UpdatePositionAndRotationPacket_1_21) GetDeltaZ() int16 {
	return p.DeltaZ
}

func (p *UpdatePositionAndRotationPacket_1_21) SetDeltaZ(deltaZ int16) {
	p.DeltaZ = deltaZ
}

func (p *UpdatePositionAndRotationPacket_1_21) GetYaw() byte {
	return p.Yaw
}

func (p *UpdatePositionAndRotationPacket_1_21) SetYaw(yaw byte) {
	p.Yaw = yaw
}

func (p *UpdatePositionAndRotationPacket_1_21) GetPitch() byte {
	return p.Pitch
}

func (p *UpdatePositionAndRotationPacket_1_21) SetPitch(pitch byte) {
	p.Pitch = pitch
}
func (p *UpdatePositionAndRotationPacket_1_21) GetOnGround() bool {
	return p.OnGround
}

func (p *UpdatePositionAndRotationPacket_1_21) SetOnGround(onGround bool) {
	p.OnGround = onGround
}

func (p *UpdatePositionAndRotationPacket_1_21) GetProtocolVersion() int {
	return p.ProtocolVersion
}

func (p *UpdatePositionAndRotationPacket_1_21) GetPacketID() int {
	return p.PacketID
}

func (p *UpdatePositionAndRotationPacket_1_21) GetState() player.ClientState {
	return p.State
}

func (p *UpdatePositionAndRotationPacket_1_21) IsClientBound() bool {
	return p.ClientBoundPacket
}
func NewUpdatePositionAndRotationPacket_1_21(entityID int32, deltaX int16, deltaY int16, deltaZ int16, yaw byte, pitch byte, onGround bool) UpdatePositionAndRotationPacket_1_21 {
	return UpdatePositionAndRotationPacket_1_21{
		PacketID:          0x2F, // Packet ID for Update Entity Position and Rotation
		ClientBoundPacket: true,
		ProtocolVersion:   767, // Or your specific protocol version
		State:             player.Play,
		EntityID:          entityID,
		DeltaX:            deltaX,
		DeltaY:            deltaY,
		DeltaZ:            deltaZ,
		Yaw:               yaw,
		Pitch:             pitch,
		OnGround:          onGround,
	}
}

func ReadUpdatePositionAndRotationPacket_1_21(packet *packet_utils.PacketReader) (*UpdatePositionAndRotationPacket_1_21, bool) {
	entityID, err := packet.ReadVarInt()
	if err != nil {
		return &UpdatePositionAndRotationPacket_1_21{}, false
	}
	deltaX, err := packet.ReadShort()
	if err != nil {
		return &UpdatePositionAndRotationPacket_1_21{}, false
	}
	deltaY, err := packet.ReadShort()
	if err != nil {
		return &UpdatePositionAndRotationPacket_1_21{}, false
	}
	deltaZ, err := packet.ReadShort()
	if err != nil {
		return &UpdatePositionAndRotationPacket_1_21{}, false
	}
	yaw, err := packet.ReadAngle()
	if err != nil {
		return &UpdatePositionAndRotationPacket_1_21{}, false
	}
	pitch, err := packet.ReadAngle()
	if err != nil {
		return &UpdatePositionAndRotationPacket_1_21{}, false
	}
	onGround, err := packet.ReadBoolean()
	if err != nil {
		return &UpdatePositionAndRotationPacket_1_21{}, false
	}
	return &UpdatePositionAndRotationPacket_1_21{
		PacketID:          0x2F,
		ClientBoundPacket: true,
		ProtocolVersion:   767,
		State:             player.Play,
		EntityID:          entityID,
		DeltaX:            deltaX,
		DeltaY:            deltaY,
		DeltaZ:            deltaZ,
		Yaw:               yaw,
		Pitch:             pitch,
		OnGround:          onGround,
	}, true
}

func (p UpdatePositionAndRotationPacket_1_21) GetPacket(player player.Player) *packets.Packet {
	//packet := packet_utils.NewPacketWriter()
	packet := player.GetPacketWritter()
	packet.Reset()
	packet.WriteVarInt(int32(p.PacketID))
	packet.WriteVarInt(p.EntityID)
	packet.WriteShort(p.DeltaX)
	packet.WriteShort(p.DeltaY)
	packet.WriteShort(p.DeltaZ)
	packet.WriteAngle(p.Yaw)
	packet.WriteAngle(p.Pitch)
	packet.WriteBoolean(p.OnGround)

	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"UpdatePositionAndRotationPacket_1_21",
		packet.GetPacketBuffer(),
		p.ClientBoundPacket,
		p.State)
	return real_packet
}
