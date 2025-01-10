package packets

import (
	"HexaUtils/entities/player"
	packetutils "HexaUtils/packets/utils"
	config "HexaUtils/server/config"
	"fmt"
)

type PlayerPacketReader_1_21 struct {
}

func NewPlayerPacketReader_1_21() *PlayerPacketReader_1_21 {
	return &PlayerPacketReader_1_21{}
}

// Cambiar readPacket a ReadPacket (con R mayúscula)
func (reader *PlayerPacketReader_1_21) ReadPacket(p player.Player, pack *packetutils.PacketReader, server_config *config.ServerConfig) {
	length, err := pack.ReadVarInt()
	if err != nil {
		fmt.Println("Error reading packet length:", err)
		return
	}
	packet_id, err := pack.ReadVarInt()
	if err != nil {
		fmt.Println("Error reading packet ID:", err)
		return
	}
	switch p.GetClientState() {
	case player.Status:
		ReadStatusStatePacket(*server_config, p, length, packet_id, pack)
	case player.Login:
		ReadLoginStatePacket(*server_config, p, length, packet_id, pack)
	case player.Configuration:
		ReadConfigurationStatePacket(*server_config, p, length, packet_id, pack)
	case player.Play:
		ReadPlayStatePacket(server_config, p, length, packet_id, *pack, reader)
	}
}
