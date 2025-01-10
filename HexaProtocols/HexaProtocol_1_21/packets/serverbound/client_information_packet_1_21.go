package serverbound

import (
	"HexaUtils/entities/player"
	"HexaUtils/packets"
	packet_utils "HexaUtils/packets/utils"
)

type ClientInformationPacket_1_21 struct {
	PacketID            int
	ServerBoundPacket   bool
	ProtocolVersion     int
	State               player.ClientState
	Locale              string
	ViewDistance        byte
	ChatMode            int32
	ChatColors          bool
	DisplayedSkinParts  byte
	MainHand            int32
	EnableTextFilter    bool
	AllowServerListings bool
}

func NewClientInformationPacket_1_21(locale string, viewDistance byte, chatMode int32, chatColors bool, displayedSkinParts byte, mainHand int32, enableTextFilter bool, allowServerListings bool) *ClientInformationPacket_1_21 {
	return &ClientInformationPacket_1_21{
		PacketID:            0x03,
		ServerBoundPacket:   true,
		ProtocolVersion:     767,
		State:               player.Status,
		Locale:              locale,
		ViewDistance:        viewDistance,
		ChatMode:            chatMode,
		ChatColors:          chatColors,
		DisplayedSkinParts:  displayedSkinParts,
		MainHand:            mainHand,
		EnableTextFilter:    enableTextFilter,
		AllowServerListings: allowServerListings,
	}
}

func ReadClientInformationPacket_1_21(packet *packet_utils.PacketReader) *ClientInformationPacket_1_21 {
	locale, err := packet.ReadString()
	if err != nil {
		return nil
	}
	viewDistance, err := packet.ReadByte()
	if err != nil {
		return nil
	}
	chatMode, err := packet.ReadVarInt()
	if err != nil {
		return nil
	}
	chatColors, err := packet.ReadBoolean()
	if err != nil {
		return nil
	}
	displayedSkinParts, err := packet.ReadUnsignedByte()
	if err != nil {
		return nil
	}
	mainHand, err := packet.ReadVarInt()
	if err != nil {
		return nil
	}
	enableTextFilter, err := packet.ReadBoolean()
	if err != nil {
		return nil
	}
	allowServerListings, err := packet.ReadBoolean()
	if err != nil {
		return nil
	}
	return NewClientInformationPacket_1_21(locale, viewDistance, chatMode, chatColors, displayedSkinParts, mainHand, enableTextFilter, allowServerListings)
}

func (p *ClientInformationPacket_1_21) GetLocale() string {
	return p.Locale
}

func (p *ClientInformationPacket_1_21) SetLocale(locale string) {
	p.Locale = locale
}

func (p *ClientInformationPacket_1_21) GetViewDistance() byte {
	return p.ViewDistance
}

func (p *ClientInformationPacket_1_21) SetViewDistance(viewDistance byte) {
	p.ViewDistance = viewDistance
}

func (p *ClientInformationPacket_1_21) GetChatMode() int32 {
	return p.ChatMode
}

func (p *ClientInformationPacket_1_21) SetChatMode(chatMode int32) {
	p.ChatMode = chatMode
}

func (p *ClientInformationPacket_1_21) GetChatColors() bool {
	return p.ChatColors
}

func (p *ClientInformationPacket_1_21) SetChatColors(chatColors bool) {
	p.ChatColors = chatColors
}

func (p *ClientInformationPacket_1_21) GetDisplayedSkinParts() byte {
	return p.DisplayedSkinParts
}

func (p *ClientInformationPacket_1_21) SetDisplayedSkinParts(displayedSkinParts byte) {
	p.DisplayedSkinParts = displayedSkinParts
}

func (p *ClientInformationPacket_1_21) GetMainHand() int32 {
	return p.MainHand
}

func (p *ClientInformationPacket_1_21) SetMainHand(mainHand int32) {
	p.MainHand = mainHand
}

func (p *ClientInformationPacket_1_21) GetEnableTextFilter() bool {
	return p.EnableTextFilter
}

func (p *ClientInformationPacket_1_21) SetEnableTextFilter(enableTextFilter bool) {
	p.EnableTextFilter = enableTextFilter
}

func (p *ClientInformationPacket_1_21) GetAllowServerListings() bool {
	return p.AllowServerListings
}

func (p *ClientInformationPacket_1_21) SetAllowServerListings(allowServerListings bool) {
	p.AllowServerListings = allowServerListings
}

func (p *ClientInformationPacket_1_21) GetPacket(player player.Player) *packets.Packet {
	//packet := packet_utils.NewPacketWriter()
	packet := player.GetPacketWritter()
	packet.Reset()
	packet.WriteVarInt(int32(p.PacketID))
	packet.WriteString(p.Locale)
	packet.WriteByte(p.ViewDistance)
	packet.WriteVarInt(p.ChatMode)
	packet.WriteBoolean(p.ChatColors)
	packet.WriteUnsignedByte(p.DisplayedSkinParts)
	packet.WriteVarInt(p.MainHand)
	packet.WriteBoolean(p.EnableTextFilter)
	packet.WriteBoolean(p.AllowServerListings)
	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"ClientInformationPacket",
		packet.GetPacketBuffer(),
		p.ServerBoundPacket,
		p.State)
	return real_packet
}
