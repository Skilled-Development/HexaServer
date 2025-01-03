package clientbound

import (
	"HexaUtils/entities/player"
	"HexaUtils/packets"
)

type PongResponsePacket_1_21 struct {
	PacketID          int
	ServerBoundPacket bool
	ProtocolVersion   int
	State             player.ClientState
	Timestamp         int64
}

func NewPongResponsePacket_1_21(timestamp int64) *PongResponsePacket_1_21 {
	return &PongResponsePacket_1_21{
		PacketID:          0x01,
		ServerBoundPacket: false,
		ProtocolVersion:   767,
		State:             player.Status,
		Timestamp:         timestamp,
	}
}

func (p *PongResponsePacket_1_21) GetTimestamp() int64 {
	return p.Timestamp
}

func (p *PongResponsePacket_1_21) GetPacket() *packets.Packet {
	packet := packets.NewPacketWriter()
	packet.WriteVarInt(int32(p.PacketID))
	packet.WriteLong(p.Timestamp)

	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"PongResponsePacket",
		packet.GetPacketBuffer(),
		p.ServerBoundPacket,
		p.State)
	return real_packet
}
