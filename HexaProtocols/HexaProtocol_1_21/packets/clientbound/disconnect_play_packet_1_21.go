package clientbound

import (
	"HexaUtils/entities/player"
	"HexaUtils/nbt"
	"HexaUtils/packets"
	"HexaUtils/style"
)

type DisconnectPacket_1_21 struct {
	PacketID          int
	ServerBoundPacket bool
	ProtocolVersion   int
	State             player.ClientState
	Reason            map[string]interface{} // Reason is the map for NBT now.
	Json              bool
}

func (p *DisconnectPacket_1_21) GetProtocolVersion() int {
	return p.ProtocolVersion
}

func (p *DisconnectPacket_1_21) GetPacketID() int {
	return p.PacketID
}

func (p *DisconnectPacket_1_21) GetState() player.ClientState {
	return p.State
}

func (p *DisconnectPacket_1_21) IsServerBound() bool {
	return p.ServerBoundPacket
}

// NewDisconnectPacket_1_21 creates a disconnect packet with the reason as a plain string
func NewDisconnectPacket_1_21(reason string) DisconnectPacket_1_21 {
	return DisconnectPacket_1_21{
		PacketID:          0x1D,
		ServerBoundPacket: false,
		ProtocolVersion:   767,
		State:             player.Login,
		Reason: map[string]interface{}{
			"text": style.ReplaceAllPlaceholders(reason),
		},
		Json: false,
	}
}

func (p *DisconnectPacket_1_21) GetPacket(player player.Player) *packets.Packet {
	//packet := packet_utils.NewPacketWriter()
	packet := player.GetPacketWritter()
	packet.Reset()
	packet.WriteVarInt(int32(p.PacketID))

	// Write TextComponent as NBT
	nbtCompound := nbt.NbtCompoundFromInterfaceMap(p.Reason)
	packet.WriteNBT(nbt.Nbt{Name: "", RootTag: nbtCompound})

	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"DisconnectPlayPacket",
		packet.GetPacketBuffer(),
		p.ServerBoundPacket,
		p.State)
	return real_packet
}
