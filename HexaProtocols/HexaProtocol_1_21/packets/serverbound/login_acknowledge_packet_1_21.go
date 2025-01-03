package serverbound

import (
	"HexaUtils/entities/player"
	"HexaUtils/packets"
)

type LoginAcknowledgePacket_1_21 struct {
	PacketID          int
	ServerBoundPacket bool
	ProtocolVersion   int
	State             player.ClientState
}

func NewLoginAcknowledgePacket_1_21() *LoginAcknowledgePacket_1_21 {
	return &LoginAcknowledgePacket_1_21{
		PacketID:          0x03,
		ServerBoundPacket: true,
		ProtocolVersion:   767,
		State:             player.Status,
	}
}

func (p *LoginAcknowledgePacket_1_21) GetPacket() *packets.Packet {
	packet := packets.NewPacketWriter()
	packet.WriteVarInt(int32(p.PacketID))
	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"LoginAcknowledgePacket",
		packet.GetPacketBuffer(),
		p.ServerBoundPacket,
		p.State)
	return real_packet
}
