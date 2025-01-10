package serverbound

import (
	"HexaUtils/entities/player"
	"HexaUtils/packets"
	packet_utils "HexaUtils/packets/utils"
)

type PluginMessageConfigurationPacket_1_21 struct {
	PacketID          int
	ServerBoundPacket bool
	ProtocolVersion   int
	State             player.ClientState
	ChannelIdentifier string
	Data              []byte
}

func (p *PluginMessageConfigurationPacket_1_21) GetChannelIdentifier() string {
	return p.ChannelIdentifier
}

func (p *PluginMessageConfigurationPacket_1_21) GetData() []byte {
	return p.Data
}

func NewPluginMessageConfigurationPacket_1_21(ChannelIdentifier string, data []byte) *PluginMessageConfigurationPacket_1_21 {
	return &PluginMessageConfigurationPacket_1_21{
		PacketID:          0x03,
		ServerBoundPacket: true,
		ProtocolVersion:   767,
		State:             player.Status,
		ChannelIdentifier: ChannelIdentifier,
		Data:              data,
	}
}
func ReadPluginMessageConfigurationPacket_1_21(packet *packet_utils.PacketReader) *PluginMessageConfigurationPacket_1_21 {
	channel_identifier, err := packet.ReadIdentifier()
	if err != nil {
		return nil
	}
	availableBytes := packet.AvailableBytes()
	if availableBytes > 32767 {
		availableBytes = 32767
	}
	data, err := packet.ReadBytes(availableBytes)
	if err != nil {
	}
	return NewPluginMessageConfigurationPacket_1_21(channel_identifier, data)
}
func (p *PluginMessageConfigurationPacket_1_21) GetPacket(player player.Player) *packets.Packet {
	//packet := packet_utils.NewPacketWriter()
	packet := player.GetPacketWritter()
	packet.Reset()
	packet.WriteVarInt(int32(p.PacketID))
	packet.WriteString(p.ChannelIdentifier)
	packet.WriteByteArray(p.Data)
	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"LoginAcknowledgePacket",
		packet.GetPacketBuffer(),
		p.ServerBoundPacket,
		p.State)
	return real_packet
}
