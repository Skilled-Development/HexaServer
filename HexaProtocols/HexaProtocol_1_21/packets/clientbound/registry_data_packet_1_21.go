package clientbound

import (
	player "HexaUtils/entities/player"
	"HexaUtils/packets"
	registries "HexaUtils/registries"
	"log"
	"os"
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

func (p *RegistryDataPacket_1_21) GetPacket() *packets.Packet {
	packet := packets.NewPacketWriter()
	packet.WriteVarInt(int32(p.PacketID))
	packet.WriteIdentifier("minecraft:" + p.Registry.GetName())
	packet.WriteVarInt(int32(len(p.GetRegistry().GetEntriesAsNBTs())))
	for _, registry := range p.GetRegistry().GetEntriesAsNBTs() {
		packet.WriteIdentifier("minecraft:" + registry.GetName())
		packet.WriteBoolean(true)
		packet.WriteNBT(registry)
	}
	packetBytes := packet.GetAsPacket().GetPacketBuffer()
	if err := os.Remove("packet_data.bin"); err != nil && !os.IsNotExist(err) {
		log.Fatalf("Error deleting packet_data.bin: %v", err)
	}
	newerror := os.WriteFile("packet_data.bin", packetBytes, 0644)
	if newerror != nil {
		log.Fatalf("Error al guardar packetBytes en el archivo: %v", newerror)
	}

	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"RegistryDataPacket",
		packet.GetPacketBuffer(),
		p.ServerBoundPacket,
		p.State)
	return real_packet
}
