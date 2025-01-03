package packets

import (
	"HexaProtocol_1_21/packets/clientbound"
	serverbound "HexaProtocol_1_21/packets/serverbound"
	"HexaUtils/entities/player"
	"HexaUtils/packets"
	config "HexaUtils/server/config"
	"fmt"
)

func ReadLoginStatePacket(server_config config.ServerConfig, p player.Player, length int32, packet_id int32, pack *packets.PacketReader) {
	switch packet_id {
	case 0x00:
		handle_login_start_packet(p, pack)
	case 0x03:
		handle_login_acknowledged_packet(p)
	default:
		fmt.Println("Unknown packet ID in Login state: ", packet_id)
	}
}

func handle_login_acknowledged_packet(p player.Player) {
	p.SetClientState(player.Configuration)
}

func handle_login_start_packet(p player.Player, pack *packets.PacketReader) {
	login_start_packet := serverbound.ReadLoginStartPacket_1_21(pack)
	if login_start_packet == nil {
		return
	}
	username := login_start_packet.GetUsername()
	fmt.Println("Username: ", username)
	uuid := login_start_packet.GetUUID()
	fmt.Println("UUID: ", uuid)
	p.SetName(username)
	p.SetUUID(uuid)

	//TODO: IMPLMENT LOGIN

	//login success packet
	login_sucess_packet := clientbound.NewLoginSuccessPacket_1_21(uuid, username, false)
	login_sucess_packet.GetPacket().Send(p)

}
