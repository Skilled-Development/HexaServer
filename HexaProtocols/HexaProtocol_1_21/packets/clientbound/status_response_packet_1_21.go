package clientbound

import (
	"HexaUtils/entities/player"
	"HexaUtils/packets"
	"HexaUtils/server/config"
)

type StatusResponsePacket struct {
	PacketID          int
	ServerBoundPacket bool
	ProtocolVersion   int
	State             player.ClientState
	Motd              config.MOTD
}

func NewStatusResponsePacket_1_21(motd *config.MOTD) *StatusResponsePacket {
	return &StatusResponsePacket{
		PacketID:          0x00,
		ServerBoundPacket: false,
		ProtocolVersion:   767,
		State:             player.Status,
		Motd:              *motd,
	}
}

func (p *StatusResponsePacket) GetMotd() config.MOTD {
	return p.Motd
}

func (p *StatusResponsePacket) SetMotd(motd config.MOTD) {
	p.Motd = motd
}

func (p *StatusResponsePacket) GetPacket() *packets.Packet {
	packet := packets.NewPacketWriter()
	packet.WriteVarInt(0x00)
	packet.WriteJson((&p.Motd).GetJSON())

	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"StatusResponsePacket",
		packet.GetPacketBuffer(),
		p.ServerBoundPacket,
		p.State)

	return real_packet
}
