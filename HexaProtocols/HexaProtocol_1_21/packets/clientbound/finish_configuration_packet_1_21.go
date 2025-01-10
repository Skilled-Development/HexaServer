package clientbound

import (
	player "HexaUtils/entities/player"
	"HexaUtils/packets"
	packet_utils "HexaUtils/packets/utils"

	"github.com/google/uuid"
)

type FinishConfigurationPacket_1_21 struct {
	PacketID          int
	ServerBoundPacket bool
	ProtocolVersion   int
	State             player.ClientState
	Writer            *packet_utils.PacketWriter
}

func NewFinishConfigurationPacket_1_21(uuid uuid.UUID, username string, errrorhandling bool) *FinishConfigurationPacket_1_21 {
	return &FinishConfigurationPacket_1_21{
		PacketID:          0x03,
		ServerBoundPacket: false,
		ProtocolVersion:   767,
		State:             player.Configuration,
	}
}
func (p *FinishConfigurationPacket_1_21) GetPacket(player player.Player) *packets.Packet {
	//packet := packet_utils.NewPacketWriter()
	packet := player.GetPacketWritter()
	packet.Reset()
	packet.WriteVarInt(int32(p.PacketID))
	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"FinishConfigurationPacket",
		packet.GetPacketBuffer(),
		p.ServerBoundPacket,
		p.State)
	return real_packet
}
