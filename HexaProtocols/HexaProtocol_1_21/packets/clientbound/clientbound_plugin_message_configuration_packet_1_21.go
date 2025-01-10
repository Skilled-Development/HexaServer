package clientbound

import (
	"HexaUtils/entities/player"
	"HexaUtils/packets"
)

type ClientboundPluginMessagePacket_1_21 struct {
	PacketID          int
	ServerBoundPacket bool
	ProtocolVersion   int
	State             player.ClientState
	ChannelIdentifier string
	Data              []byte
}

func NewClientboundPluginMessagePacket_1_21(channel string, data []byte) *ClientboundPluginMessagePacket_1_21 {
	return &ClientboundPluginMessagePacket_1_21{
		PacketID:          0x01,
		ServerBoundPacket: false,
		ProtocolVersion:   767,
		State:             player.Status,
		ChannelIdentifier: channel,
		Data:              data,
	}
}

func (p *ClientboundPluginMessagePacket_1_21) GetChannelIdentifier() string {
	return p.ChannelIdentifier
}

func (p *ClientboundPluginMessagePacket_1_21) SetChannelIdentifier(channel string) {
	p.ChannelIdentifier = channel
}

func (p *ClientboundPluginMessagePacket_1_21) GetData() []byte {
	return p.Data
}

func (p *ClientboundPluginMessagePacket_1_21) SetData(data []byte) {
	p.Data = data
}

func (p *ClientboundPluginMessagePacket_1_21) GetPacket(player player.Player) *packets.Packet {
	//packet := packet_utils.NewPacketWriter()
	packet := player.GetPacketWritter()
	packet.Reset()
	packet.WriteVarInt(int32(p.PacketID))
	packet.WriteString(p.ChannelIdentifier)
	packet.WriteByteArray(p.Data)
	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"ClientboundPluginMessagePacket",
		packet.GetPacketBuffer(),
		p.ServerBoundPacket,
		p.State)
	return real_packet
}
