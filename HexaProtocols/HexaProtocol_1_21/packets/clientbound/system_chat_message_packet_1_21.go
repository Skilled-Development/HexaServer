package clientbound

import (
	"HexaUtils/entities/player"
	"HexaUtils/nbt"
	"HexaUtils/packets"
)

type SystemChatMessagePacket_1_21 struct {
	PacketID          int
	ServerBoundPacket bool
	ProtocolVersion   int
	State             player.ClientState
	Content           string
	Overlay           bool
}

func NewSystemChatMessagePacket_1_21(content string, overlay bool) *SystemChatMessagePacket_1_21 {
	return &SystemChatMessagePacket_1_21{
		PacketID:          0x6C,
		ServerBoundPacket: false,
		ProtocolVersion:   767, // Assuming protocol version is the same
		State:             player.Play,
		Content:           content,
		Overlay:           overlay,
	}
}

func (p *SystemChatMessagePacket_1_21) GetContent() string {
	return p.Content
}

func (p *SystemChatMessagePacket_1_21) SetContent(content string) {
	p.Content = content
}
func (p *SystemChatMessagePacket_1_21) GetOverlay() bool {
	return p.Overlay
}

func (p *SystemChatMessagePacket_1_21) SetOverlay(overlay bool) {
	p.Overlay = overlay
}

func (p *SystemChatMessagePacket_1_21) GetPacket(player player.Player) *packets.Packet {
	//packet := packet_utils.NewPacketWriter()
	packet := player.GetPacketWritter()
	packet.Reset()
	packet.WriteVarInt(int32(p.PacketID))
	compound := nbt.NbtCompoundFromInterfaceMap(map[string]interface{}{
		"type": "text",
		"text": p.Content,
	})
	nbt := nbt.NewNbt("TEXT", compound)
	packet.WriteNBT(*nbt)

	packet.WriteBoolean(p.Overlay)

	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"SystemChatMessagePacket",
		packet.GetPacketBuffer(),
		p.ServerBoundPacket,
		p.State)
	return real_packet
}
