package clientbound

import (
	player "HexaUtils/entities/player"
	"HexaUtils/packets"
	packet_utils "HexaUtils/packets/utils"

	"github.com/google/uuid"
)

type LoginSuccessPacket_1_21 struct {
	PacketID          int
	ServerBoundPacket bool
	ProtocolVersion   int
	State             player.ClientState
	UUID              uuid.UUID
	Username          string
	//TODO: add properties
	StrictErrorHandling bool
	Writer              *packet_utils.PacketWriter
}

func NewLoginSuccessPacket_1_21(uuid uuid.UUID, username string, errrorhandling bool) *LoginSuccessPacket_1_21 {
	return &LoginSuccessPacket_1_21{
		PacketID:          0x02,
		ServerBoundPacket: false,
		ProtocolVersion:   767,
		State:             player.Login,
		UUID:              uuid,
		Username:          username,
		//TODO: add properties
		StrictErrorHandling: errrorhandling,
	}
}
func (p *LoginSuccessPacket_1_21) GetPacket(player player.Player) *packets.Packet {
	//packet := packet_utils.NewPacketWriter()
	packet := player.GetPacketWritter()
	packet.Reset()
	packet.WriteVarInt(int32(p.PacketID))
	packet.WriteUUID(p.UUID)
	packet.WriteString(p.Username)
	//TODO: add properties
	packet.WriteVarInt(0) //properties length
	packet.WriteBoolean(p.StrictErrorHandling)
	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"LoginSuccessPacket",
		packet.GetPacketBuffer(),
		p.ServerBoundPacket,
		p.State)
	return real_packet
}
