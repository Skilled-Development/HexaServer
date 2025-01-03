package serverbound

import (
	"HexaUtils/entities/player"
	"HexaUtils/packets"

	"github.com/google/uuid"
)

type LoginStartPacket_1_21 struct {
	PacketID          int
	ServerBoundPacket bool
	ProtocolVersion   int
	State             player.ClientState
	Username          string
	UUID              uuid.UUID
}

func NewLoginStartPacket_1_21(name string, uuid uuid.UUID) *LoginStartPacket_1_21 {
	return &LoginStartPacket_1_21{
		PacketID:          0x01,
		ServerBoundPacket: true,
		ProtocolVersion:   767,
		State:             player.Login,
		Username:          name,
		UUID:              uuid,
	}
}

func ReadLoginStartPacket_1_21(packet *packets.PacketReader) *LoginStartPacket_1_21 {
	name, err := packet.ReadString()
	if err != nil {
		return nil
	}
	uuid, err := packet.ReadUUID()
	if err != nil {
		return nil
	}
	return NewLoginStartPacket_1_21(name, uuid)
}

func (p *LoginStartPacket_1_21) GetUsername() string {
	return p.Username
}

func (p *LoginStartPacket_1_21) GetUUID() uuid.UUID {
	return p.UUID
}

func (p *LoginStartPacket_1_21) GetPacket() *packets.Packet {
	packet := packets.NewPacketWriter()
	packet.WriteVarInt(int32(p.PacketID))
	packet.WriteString(p.Username)
	packet.WriteUUID(p.UUID)

	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"LoginStartPacket",
		packet.GetPacketBuffer(),
		p.ServerBoundPacket,
		p.State)
	return real_packet
}
