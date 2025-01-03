package serverbound

import (
	player "HexaUtils/entities/player"
	"HexaUtils/packets"
)

type KeepAlivePlayPacket_1_21 struct {
	PacketID          int
	ServerBoundPacket bool
	ProtocolVersion   int
	State             player.ClientState
	KeepAliveID       int64
}

func (p *KeepAlivePlayPacket_1_21) GetKeepAliveID() int64 {
	return p.KeepAliveID
}

func (p *KeepAlivePlayPacket_1_21) SetKeepAliveID(keepAliveID int64) {
	p.KeepAliveID = keepAliveID
}

func NewKeepAlivePlayPacket_1_21(keepAliveID int64) *KeepAlivePlayPacket_1_21 {
	return &KeepAlivePlayPacket_1_21{
		PacketID:          0x18,
		ServerBoundPacket: false,
		ProtocolVersion:   767,
		State:             player.Play,
		KeepAliveID:       keepAliveID,
	}
}

func ReadKeepAlivePlayPacket_1_21(packet packets.PacketReader) *KeepAlivePlayPacket_1_21 {
	keepAliveID, err := packet.ReadLong()
	if err != nil {
		return nil
	}
	return &KeepAlivePlayPacket_1_21{
		PacketID:          0x1A,
		ServerBoundPacket: true,
		ProtocolVersion:   767,
		State:             player.Play,
		KeepAliveID:       keepAliveID,
	}
}

func (p *KeepAlivePlayPacket_1_21) GetProtocolVersion() int {
	return p.ProtocolVersion
}

func (p *KeepAlivePlayPacket_1_21) GetPacketID() int {
	return p.PacketID
}

func (p *KeepAlivePlayPacket_1_21) GetState() player.ClientState {
	return p.State
}

func (p *KeepAlivePlayPacket_1_21) IsServerBound() bool {
	return p.ServerBoundPacket
}

func (p KeepAlivePlayPacket_1_21) GetPacket() *packets.Packet {
	packet := packets.NewPacketWriter()
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
