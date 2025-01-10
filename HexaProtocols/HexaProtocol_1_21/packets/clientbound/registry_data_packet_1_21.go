package clientbound

import (
	player "HexaUtils/entities/player"
	"HexaUtils/packets"
	registries "HexaUtils/registries"
)

type RegistryDataPacket_1_21 struct {
	PacketID          int
	ServerBoundPacket bool
	ProtocolVersion   int
	State             player.ClientState
	Registry          registries.Registry
}

func NewRegistryDataPacket_1_21(registry registries.Registry) *RegistryDataPacket_1_21 {
	return &RegistryDataPacket_1_21{
		PacketID:          0x07,
		ServerBoundPacket: false,
		ProtocolVersion:   767,
		State:             player.Login,
		Registry:          registry,
	}
}

func (p *RegistryDataPacket_1_21) GetRegistry() registries.Registry {
	return p.Registry
}

func (p *RegistryDataPacket_1_21) GetPacket(player player.Player) *packets.Packet {
	//packet := packet_utils.NewPacketWriter()
	packet := player.GetPacketWritter()
	packet.Reset()
	packet.WriteVarInt(int32(p.PacketID))
	packet.WriteIdentifier("minecraft:" + p.Registry.GetName())
	packet.WriteVarInt(int32(len(p.GetRegistry().GetEntriesAsNBTs())))
	for _, registry := range p.GetRegistry().GetEntriesAsNBTs() {
		packet.WriteIdentifier("minecraft:" + registry.GetName())
		packet.WriteBoolean(true)
		packet.WriteNBT(registry)
	}

	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"RegistryDataPacket",
		packet.GetPacketBuffer(),
		p.ServerBoundPacket,
		p.State)
	return real_packet
}
