package clientbound

import (
	"HexaUtils/entities/player"
	"HexaUtils/packets"
)

type SetCenterChunkPacket_1_21 struct {
	PacketID          int
	ServerBoundPacket bool
	ProtocolVersion   int
	State             player.ClientState
	ChunkX            int32
	ChunkZ            int32
}

func NewSetCenterChunkPacket_1_21(chunkX int32, chunkZ int32) *SetCenterChunkPacket_1_21 {
	return &SetCenterChunkPacket_1_21{
		PacketID:          0x54,
		ServerBoundPacket: false,
		ProtocolVersion:   767,
		State:             player.Play,
		ChunkX:            chunkX,
		ChunkZ:            chunkZ,
	}
}

func (p *SetCenterChunkPacket_1_21) GetChunkX() int32 {
	return p.ChunkX
}

func (p *SetCenterChunkPacket_1_21) SetChunkX(chunkX int32) {
	p.ChunkX = chunkX
}

func (p *SetCenterChunkPacket_1_21) GetChunkZ() int32 {
	return p.ChunkZ
}

func (p *SetCenterChunkPacket_1_21) SetChunkZ(chunkZ int32) {
	p.ChunkZ = chunkZ
}

func (p *SetCenterChunkPacket_1_21) GetPacket(player player.Player) *packets.Packet {
	//packet := packet_utils.NewPacketWriter()
	packet := player.GetPacketWritter()
	packet.Reset()
	packet.WriteVarInt(int32(p.PacketID))
	packet.WriteVarInt(p.ChunkX)
	packet.WriteVarInt(p.ChunkZ)
	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"SetCenterChunkPacket",
		packet.GetPacketBuffer(),
		p.ServerBoundPacket,
		p.State)
	return real_packet
}
