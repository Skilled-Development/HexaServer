package clientbound

import (
	player "HexaUtils/entities/player"
	"HexaUtils/packets"
	packet_utils "HexaUtils/packets/utils"
)

type ClientboundKnownPacks_1_21 struct {
	PacketID          int
	ServerBoundPacket bool
	ProtocolVersion   int
	State             player.ClientState
	KnownPacksCount   int
	KnownPacks        []packet_utils.KnownPack
}

func (p *ClientboundKnownPacks_1_21) GetKnownPacksCount() int {
	return p.KnownPacksCount
}

func (p *ClientboundKnownPacks_1_21) SetKnownPacksCount(knownPacksCount int) {
	p.KnownPacksCount = knownPacksCount
}

func (p *ClientboundKnownPacks_1_21) GetKnownPacks() []packet_utils.KnownPack {
	return p.KnownPacks
}

func (p *ClientboundKnownPacks_1_21) SetKnownPacks(knownPacks []packet_utils.KnownPack) {
	p.KnownPacks = knownPacks
}

func NewClientboundKnownPacks_1_21(packs []packet_utils.KnownPack) *ClientboundKnownPacks_1_21 {
	return &ClientboundKnownPacks_1_21{
		PacketID:          0x0E,
		ServerBoundPacket: false,
		ProtocolVersion:   767,
		State:             player.Login,
		KnownPacksCount:   len(packs),
		KnownPacks:        packs,
	}
}
func (p *ClientboundKnownPacks_1_21) GetPacket(player player.Player) *packets.Packet {
	//packet := packet_utils.NewPacketWriter()
	packet := player.GetPacketWritter()
	packet.Reset()
	packet.WriteVarInt(int32(p.PacketID))
	packet.WriteVarInt(int32(p.KnownPacksCount))
	for _, pack := range p.KnownPacks {
		packet.WriteString(pack.GetNamespace())
		packet.WriteString(pack.GetID())
		packet.WriteString(pack.GetVersion())
	}
	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"ClientboundKnownPacksPacket",
		packet.GetPacketBuffer(),
		p.ServerBoundPacket,
		p.State)
	return real_packet
}
