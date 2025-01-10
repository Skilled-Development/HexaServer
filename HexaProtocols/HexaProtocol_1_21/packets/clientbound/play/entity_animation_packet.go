package play

import (
	"HexaUtils/entities/player"
	"HexaUtils/packets"
	packet_utils "HexaUtils/packets/utils"
)

type AnimationID int

const (
	SwingMainArm        AnimationID = 0
	LeaveBed            AnimationID = 2
	SwingOffhand        AnimationID = 3
	CriticalEffect      AnimationID = 4
	MagicCriticalEffect AnimationID = 5
)

type EntityAnimationPacket_1_21 struct {
	PacketID          int
	ClientBoundPacket bool
	ProtocolVersion   int
	State             player.ClientState
	EntityID          int32
	Animation         AnimationID
}

func (p *EntityAnimationPacket_1_21) GetEntityID() int32 {
	return p.EntityID
}

func (p *EntityAnimationPacket_1_21) SetEntityID(entityID int32) {
	p.EntityID = entityID
}

func (p *EntityAnimationPacket_1_21) GetAnimation() AnimationID {
	return p.Animation
}

func (p *EntityAnimationPacket_1_21) SetAnimation(animation AnimationID) {
	p.Animation = animation
}

func (p *EntityAnimationPacket_1_21) GetProtocolVersion() int {
	return p.ProtocolVersion
}

func (p *EntityAnimationPacket_1_21) GetPacketID() int {
	return p.PacketID
}

func (p *EntityAnimationPacket_1_21) GetState() player.ClientState {
	return p.State
}

func (p *EntityAnimationPacket_1_21) IsClientBound() bool {
	return p.ClientBoundPacket
}

func NewEntityAnimationPacket_1_21(entityID int32, animation AnimationID) EntityAnimationPacket_1_21 {
	return EntityAnimationPacket_1_21{
		PacketID:          0x03,
		ClientBoundPacket: true,
		ProtocolVersion:   767, // Or your specific protocol version
		State:             player.Play,
		EntityID:          entityID,
		Animation:         animation,
	}
}

func ReadEntityAnimationPacket_1_21(packet *packet_utils.PacketReader) (*EntityAnimationPacket_1_21, bool) {
	entityID, err := packet.ReadVarInt()
	if err != nil {
		return &EntityAnimationPacket_1_21{}, false
	}
	animation, err := packet.ReadUnsignedByte()
	if err != nil {
		return &EntityAnimationPacket_1_21{}, false
	}
	return &EntityAnimationPacket_1_21{
		PacketID:          0x03,
		ClientBoundPacket: true,
		ProtocolVersion:   767,
		State:             player.Play,
		EntityID:          entityID,
		Animation:         AnimationID(animation),
	}, true
}

func (p EntityAnimationPacket_1_21) GetPacket(player player.Player) *packets.Packet {
	//packet := packet_utils.NewPacketWriter()
	packet := player.GetPacketWritter()
	packet.Reset()
	packet.WriteVarInt(int32(p.PacketID))
	packet.WriteVarInt(p.EntityID)
	packet.WriteUnsignedByte(uint8(p.Animation))

	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"EntityAnimationPacket",
		packet.GetPacketBuffer(),
		p.ClientBoundPacket,
		p.State)
	return real_packet
}
