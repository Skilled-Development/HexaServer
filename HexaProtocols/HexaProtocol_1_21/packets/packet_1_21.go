package packets

import (
	"HexaUtils/entities/player"
	"HexaUtils/packets"
)

type Packet_1_21 interface {
	//GetPacket() *packets.Packet
	GetPacket(p player.Player) *packets.Packet
}

type QueuedPacket struct {
	Player     player.Player
	PacketData Packet_1_21
}
