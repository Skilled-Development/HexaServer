package packets

import (
	clientbound "HexaProtocol_1_21/packets/clientbound"
	serverbound "HexaProtocol_1_21/packets/serverbound"
	"HexaUtils/entities/player"
	packet_utils "HexaUtils/packets/utils"
	config "HexaUtils/server/config"
	"fmt"
)

func ReadStatusStatePacket(server_config config.ServerConfig, p player.Player, length int32, packet_id int32, pack *packet_utils.PacketReader) {
	switch packet_id {
	case 0x00:
		handle_status_request_packet(p, server_config)
	case 0x01:
		handle_ping_request_packet(p, pack)
	default:
		fmt.Println("Unknown packet ID in Status state: ", packet_id)
	}
}

func handle_ping_request_packet(p player.Player, pack *packet_utils.PacketReader) {
	ping_response_packet := serverbound.ReadPingRequestPacket_1_21(pack)
	if ping_response_packet == nil {
		return
	}
	timestamp := ping_response_packet.GetTimestamp()
	clientbound.NewPongResponsePacket_1_21(timestamp).GetPacket(p).Send(p)
}

func handle_status_request_packet(p player.Player, server_config config.ServerConfig) {
	motd := server_config.GetMOTD()
	clientbound.NewStatusResponsePacket_1_21(&motd).GetPacket(p).Send(p)
}
