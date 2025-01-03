package serverbound

import (
	player "HexaUtils/entities/player"
	"HexaUtils/packets"
	packetutils "HexaUtils/packets/utils"
)

type ServerboundKnownPacks_1_21 struct {
	PacketID          int
	ServerBoundPacket bool
	ProtocolVersion   int
	State             player.ClientState
	KnownPacksCount   int
	KnownPacks        []packetutils.KnownPack
}

func (p *ServerboundKnownPacks_1_21) GetKnownPacksCount() int {
	return p.KnownPacksCount
}

func (p *ServerboundKnownPacks_1_21) SetKnownPacksCount(knownPacksCount int) {
	p.KnownPacksCount = knownPacksCount
}

func (p *ServerboundKnownPacks_1_21) GetKnownPacks() []packetutils.KnownPack {
	return p.KnownPacks
}

func (p *ServerboundKnownPacks_1_21) SetKnownPacks(knownPacks []packetutils.KnownPack) {
	p.KnownPacks = knownPacks
}

func NewServerboundKnownPacks_1_21(packs []packetutils.KnownPack) *ServerboundKnownPacks_1_21 {
	return &ServerboundKnownPacks_1_21{
		PacketID:          0x07,
		ServerBoundPacket: false,
		ProtocolVersion:   767,
		State:             player.Login,
		KnownPacksCount:   len(packs),
		KnownPacks:        packs,
	}
}

func ReadServerboundKnownPacks_1_21(packet *packets.PacketReader) *ServerboundKnownPacks_1_21 {
	knownPacksCount, err := packet.ReadVarInt()
	if err != nil {
		return nil
	}
	knownPacks := make([]packetutils.KnownPack, knownPacksCount)
	for i := 0; i < int(knownPacksCount); i++ {
		namespace, err := packet.ReadString()
		if err != nil {
			return nil
		}
		id, err := packet.ReadString()
		if err != nil {
			return nil
		}
		version, err := packet.ReadString()
		if err != nil {
			return nil
		}
		knownPacks[i] = *packetutils.NewKnownPack(namespace, id, version)
	}
	return &ServerboundKnownPacks_1_21{
		PacketID:          0x07,
		ServerBoundPacket: true,
		ProtocolVersion:   767,
		State:             player.Login,
		KnownPacksCount:   int(knownPacksCount),
		KnownPacks:        knownPacks,
	}
}

func (p *ServerboundKnownPacks_1_21) GetPacket() *packets.Packet {
	packet := packets.NewPacketWriter()
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
