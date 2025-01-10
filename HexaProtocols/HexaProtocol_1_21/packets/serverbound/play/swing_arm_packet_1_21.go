// serverbound/swingarm.go
package serverbound

import (
	player "HexaUtils/entities/player"
	"HexaUtils/packets"
	packet_utils "HexaUtils/packets/utils"
)

type SwingArmPacket_1_21 struct {
	PacketID          int
	ServerBoundPacket bool
	ProtocolVersion   int
	State             player.ClientState
	Hand              player.Hand
}

func (p *SwingArmPacket_1_21) GetHand() player.Hand {
	return p.Hand
}

func (p *SwingArmPacket_1_21) SetHand(hand player.Hand) {
	p.Hand = hand
}

func (p *SwingArmPacket_1_21) GetProtocolVersion() int {
	return p.ProtocolVersion
}

func (p *SwingArmPacket_1_21) GetPacketID() int {
	return p.PacketID
}

func (p *SwingArmPacket_1_21) GetState() player.ClientState {
	return p.State
}

func (p *SwingArmPacket_1_21) IsServerBound() bool {
	return p.ServerBoundPacket
}

func NewSwingArmPacket_1_21(hand player.Hand) SwingArmPacket_1_21 {
	return SwingArmPacket_1_21{
		PacketID:          0x36,
		ServerBoundPacket: true,
		ProtocolVersion:   767,
		State:             player.Play,
		Hand:              hand,
	}
}

func ReadSwingArmPacket_1_21(packet packet_utils.PacketReader) (SwingArmPacket_1_21, bool) {
	hand, err := packet.ReadVarInt()
	if err != nil {
		return SwingArmPacket_1_21{}, false
	}

	return SwingArmPacket_1_21{
		PacketID:          0x36,
		ServerBoundPacket: true,
		ProtocolVersion:   767,
		State:             player.Play,
		Hand:              player.Hand(hand),
	}, true
}

func (p SwingArmPacket_1_21) GetPacket(player player.Player) *packets.Packet {
	//packet := packet_utils.NewPacketWriter()
	packet := player.GetPacketWritter()
	packet.Reset()
	packet.WriteVarInt(int32(p.PacketID))
	packet.WriteVarInt(int32(p.Hand))

	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"SwingArmPacket",
		packet.GetPacketBuffer(),
		p.ServerBoundPacket,
		p.State)
	return real_packet
}
