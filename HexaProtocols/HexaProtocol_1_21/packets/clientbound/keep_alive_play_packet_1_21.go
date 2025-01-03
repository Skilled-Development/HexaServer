package clientbound

import (
	"HexaUtils/entities/player"
	"HexaUtils/packets"
	"fmt"
)

type KeepAlivePlayPacket_1_21 struct {
	PacketID          int
	ServerBoundPacket bool
	ProtocolVersion   int
	State             player.ClientState
	KeepAliveID       int64
}

func NewKeepAlivePlayPacket_1_21(keepAliveID int64) KeepAlivePlayPacket_1_21 {
	return KeepAlivePlayPacket_1_21{
		PacketID:          0x26,
		ServerBoundPacket: false,
		ProtocolVersion:   767,
		State:             player.Play,
		KeepAliveID:       keepAliveID,
	}
}

func (p *KeepAlivePlayPacket_1_21) GetKeepAliveID() int64 {
	return p.KeepAliveID
}

func (p *KeepAlivePlayPacket_1_21) SetKeepAliveID(keepAliveID int64) {
	p.KeepAliveID = keepAliveID
}

func (p *KeepAlivePlayPacket_1_21) GetPacket() *packets.Packet {
	packet := packets.NewPacketWriter()
	println("KeepAlivePlayPacket_1_21: Writing packet ID")
	println(fmt.Sprintf("0x%X", p.PacketID))
	packet.WriteVarInt(int32(p.PacketID))
	packet.WriteLong(p.KeepAliveID)
	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"KeepAlivePlayPacket",
		packet.GetPacketBuffer(),
		p.ServerBoundPacket,
		p.State)
	return real_packet
}
