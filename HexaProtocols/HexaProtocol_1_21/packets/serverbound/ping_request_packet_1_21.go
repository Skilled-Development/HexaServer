package serverbound

import (
	"HexaUtils/entities/player"
	"HexaUtils/packets"
)

type PingRequestPacket_1_21 struct {
	PacketID          int
	ServerBoundPacket bool
	ProtocolVersion   int
	State             player.ClientState
	Timestamp         int64
}

func NewPingRequestPacket_1_21(timestamp int64) *PingRequestPacket_1_21 {
	return &PingRequestPacket_1_21{
		PacketID:          0x01,
		ServerBoundPacket: true,
		ProtocolVersion:   767,
		State:             player.Status,
		Timestamp:         timestamp,
	}
}

func ReadPingRequestPacket_1_21(packet *packets.PacketReader) *PingRequestPacket_1_21 {
	timestamp, err := packet.ReadLong()
	if err != nil {
		return nil
	}
	return NewPingRequestPacket_1_21(timestamp)
}

func (p *PingRequestPacket_1_21) GetTimestamp() int64 {
	return p.Timestamp
}

func (p *PingRequestPacket_1_21) GetPacket() *packets.Packet {
	packet := packets.NewPacketWriter()
	packet.WriteVarInt(int32(p.PacketID))
	packet.WriteLong(p.Timestamp)
	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"PingRequestPacket",
		packet.GetPacketBuffer(),
		p.ServerBoundPacket,
		p.State)
	return real_packet
}