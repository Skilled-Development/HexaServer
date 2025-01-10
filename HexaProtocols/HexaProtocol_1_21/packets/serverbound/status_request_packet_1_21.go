package serverbound

import (
	"HexaUtils/entities/player"
	"HexaUtils/packets"
)

type StatusRequestPacket_1_21 struct {
	PacketID          int
	ServerBoundPacket bool
	ProtocolVersion   int
	State             player.ClientState
}

func NewStatusRequestPacket_1_21() *StatusRequestPacket_1_21 {
	return &StatusRequestPacket_1_21{
		PacketID:          0x00,
		ServerBoundPacket: true,
		ProtocolVersion:   767,
		State:             player.Status,
	}
}

func (p *StatusRequestPacket_1_21) GetPacket(player player.Player) *packets.Packet {
	//packet := packet_utils.NewPacketWriter()
	packet := player.GetPacketWritter()
	packet.Reset()
	packet.WriteVarInt(int32(p.PacketID))
	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"StatusRequestPacket",
		packet.GetPacketBuffer(),
		p.ServerBoundPacket,
		p.State)
	return real_packet
}
