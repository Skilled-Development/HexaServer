package packets

import (
	"HexaProtocol_1_21/packets/clientbound"
	"HexaProtocol_1_21/packets/serverbound"
	entities_manager "HexaServer/entities/manager"
	"HexaUtils/entities"
	"HexaUtils/entities/player"
	"HexaUtils/packets"
	config "HexaUtils/server/config"
	"fmt"
	"time"
)

var packetQueue []QueuedPacket

func EnqueuePacket(p player.Player, pack Packet_1_21) {
	packetQueue = append(packetQueue, QueuedPacket{Player: p, PacketData: pack})
}

func ProcessPacketQueue() {
	for len(packetQueue) > 0 {
		queuePacket := packetQueue[0]
		packetQueue = packetQueue[1:]
		handlePackets(queuePacket.Player, queuePacket.PacketData)
	}
}

func handlePackets(player player.Player, packet Packet_1_21) {
	switch p := packet.(type) {
	case serverbound.ConfirmTeleportation_1_21:
		handle_confirm_teleport_packet(player, p)
	case serverbound.ServerboundPlayerPositionAndRotation_1_21:
		handle_serverbound_player_position_and_rotation_packet(player, p)
	case serverbound.KeepAlivePlayPacket_1_21:
		handle_serverbound_keep_alive_packet(player, p)
	case serverbound.ServerboundPlayerPositionPacket_1_21:
		handle_serverbound_position_packet(player, p)
	default:
		fmt.Println("Paquete desconocido recibido", packet.GetPacket().GetPacketID())
	}
}

func ReadPlayStatePacket(server_config *config.ServerConfig, p player.Player, length int32, packet_id int32, pack packets.PacketReader, reader *PlayerPacketReader_1_21) {
	println("Current packet ID:", fmt.Sprintf("0x%X", packet_id))
	println("Current packet ID:", packet_id)
	if p.NeedsKeepAlivePacket() {
		keepAlivePacket := clientbound.NewKeepAlivePlayPacket_1_21(p.GenerateKeepAliveID())
		keepAlivePacket.GetPacket().Send(p)
	}
	switch packet_id {
	case 0x00:
		create_confirm_teleport_packet(p, pack)
	case 0x1B:
		create_serverbound_player_position_and_rotation_packet(p, pack)
	case 0x18:
		create_serverbound_keep_alive_packet(p, pack)
	case 0x1A:
		create_serverbound_player_position_packet(p, pack)
	}

}

func create_confirm_teleport_packet(p player.Player, pack packets.PacketReader) {
	confirm_teleport_packet := serverbound.ReadConfirmTeleportation_1_21(pack)
	EnqueuePacket(p, confirm_teleport_packet)
}

func handle_confirm_teleport_packet(p player.Player, confirm_teleport_packet serverbound.ConfirmTeleportation_1_21) {
	teleportID := confirm_teleport_packet.GetTeleportId()
	if teleportID != p.GetTeleportID() {
		disconnectPacket := clientbound.NewDisconnectPacket_1_21(config.ServerConfigInstance.GetInvalidTpIdMessage())
		disconnectPacket.GetPacket().Send(p)
		time.Sleep(50 * time.Millisecond)
		(*p.GetConn()).Close()
		return
	}
}

func create_serverbound_player_position_and_rotation_packet(p player.Player, pack packets.PacketReader) {
	serverbound_player_position_and_rotation_packet, ok := serverbound.ReadServerboundPlayerPositionAndRotation_1_21(pack)
	if !ok {
		return
	}
	EnqueuePacket(p, serverbound_player_position_and_rotation_packet)
}
func handle_serverbound_player_position_and_rotation_packet(p player.Player, serverbound_player_position_and_rotation_packet serverbound.ServerboundPlayerPositionAndRotation_1_21) {
	last_location := p.GetLocation()
	x := serverbound_player_position_and_rotation_packet.GetX()
	feetY := serverbound_player_position_and_rotation_packet.GetFeetY()
	z := serverbound_player_position_and_rotation_packet.GetZ()
	yaw := serverbound_player_position_and_rotation_packet.GetYaw()
	pitch := serverbound_player_position_and_rotation_packet.GetPitch()
	onGround := serverbound_player_position_and_rotation_packet.GetOnGround()
	location := entities.Location{
		X:     x,
		Y:     feetY,
		Z:     z,
		Yaw:   yaw,
		Pitch: pitch,
	}
	//en vez de hacer esto solo los proceso y los guardo en la lista para luego actualizar al jugador
	p.SetLocation(&location)
	p.SetOnGround(onGround)
	allEntities := entities_manager.GetAllEntities()

	current_player_skin_signature := p.GetSkinSignature()
	property := clientbound.PlayerProperty{
		Name:      "textures",
		Value:     p.GetSkinValue(),
		IsSigned:  true,
		Signature: current_player_skin_signature,
	}
	add_player_data_action := clientbound.AddPlayerData{
		Name:       p.GetName(),
		Properties: []clientbound.PlayerProperty{property},
	}
	player_action := clientbound.PlayerActionData{
		ActionType:    clientbound.AddPlayerAction,
		AddPlayerData: &add_player_data_action,
	}
	player_info_entry := clientbound.PlayerInfoEntry{
		UUID:          p.GetUUID(),
		PlayerActions: []clientbound.PlayerActionData{player_action},
	}
	for _, entity := range allEntities {
		if entity.GetEntityType() != entities.Player {
			continue
		}
		if entity.GetEntityId() == p.GetEntityId() {
			continue
		}
		updatePacket := clientbound.NewPlayerInfoUpdatePacket_1_21(clientbound.AddPlayerAction, []clientbound.PlayerInfoEntry{player_info_entry})
		var entityAsPlayer player.Player = entity.(player.Player) // Posible optimizacion, revisar si se puede evitar

		updatePacket.GetPacket().Send(entityAsPlayer)

		current_other_skin_signature := entityAsPlayer.GetSkinSignature()
		other_property := clientbound.PlayerProperty{
			Name:      "textures",
			Value:     entityAsPlayer.GetSkinValue(),
			IsSigned:  true,
			Signature: current_other_skin_signature,
		}
		add_other_data_action := clientbound.AddPlayerData{
			Name:       entityAsPlayer.GetName(),
			Properties: []clientbound.PlayerProperty{other_property},
		}
		other_action := clientbound.PlayerActionData{
			ActionType:    clientbound.AddPlayerAction,
			AddPlayerData: &add_other_data_action,
		}
		other_info_entry := clientbound.PlayerInfoEntry{
			UUID:          p.GetUUID(),
			PlayerActions: []clientbound.PlayerActionData{other_action},
		}

		updateOtherPacket := clientbound.NewPlayerInfoUpdatePacket_1_21(clientbound.AddPlayerAction, []clientbound.PlayerInfoEntry{other_info_entry})
		updateOtherPacket.GetPacket().Send(p)
		println("Enviando PlayerInfoUpdate")
	}

	last_chunk_x := int32(last_location.X / 16)
	last_chunk_z := int32(last_location.Z / 16)
	current_chunk_x := int32(location.X / 16)
	current_chunk_z := int32(location.Z / 16)
	if last_chunk_x != current_chunk_x || last_chunk_z != current_chunk_z {
		set_center_chunk_packet := clientbound.NewSetCenterChunkPacket_1_21(current_chunk_x, current_chunk_z)
		set_center_chunk_packet.GetPacket().Send(p)
		println("Enviando SetCenterChunk")
	}
	if p.NeedsKeepAlivePacket() {
		keepAlivePacket := clientbound.NewKeepAlivePlayPacket_1_21(p.GenerateKeepAliveID())
		keepAlivePacket.GetPacket().Send(p)
		println("Enviando KeepAlive")
	}
}

func create_serverbound_keep_alive_packet(p player.Player, pack packets.PacketReader) {
	serverbound_keep_alive_packet := serverbound.ReadKeepAlivePlayPacket_1_21(pack)
	if serverbound_keep_alive_packet == nil {
		return
	}
	EnqueuePacket(p, serverbound_keep_alive_packet)
}

func handle_serverbound_keep_alive_packet(p player.Player, serverbound_keep_alive_packet serverbound.KeepAlivePlayPacket_1_21) {
	keepAliveID := serverbound_keep_alive_packet.GetKeepAliveID()
	if keepAliveID != p.GetKeepAliveID() {
		/*
			TODO: Implementar el envío de un paquete de desconexión
			disconnectPacket := clientbound.NewDisconnectPacket_1_21(config.GetInvalidKeepAliveIdMessage())
			disconnectPacket.GetPacket().Send(p)
			time.Sleep(50 * time.Millisecond)*/
		(*p.GetConn()).Close()
	}

}

func create_serverbound_player_position_packet(p player.Player, pack packets.PacketReader) {
	serverbound_player_position_packet, ok := serverbound.ReadServerboundPlayerPositionPacket_1_21(pack)
	if !ok {
		return
	}
	EnqueuePacket(p, serverbound_player_position_packet)
}

func handle_serverbound_position_packet(player player.Player, serverboundPlayerPositionPacket serverbound.ServerboundPlayerPositionPacket_1_21) {
	lastLocation := player.GetLocation()
	x := serverboundPlayerPositionPacket.GetX()
	feetY := serverboundPlayerPositionPacket.GetFeetY()
	z := serverboundPlayerPositionPacket.GetZ()
	onGround := serverboundPlayerPositionPacket.GetOnGround()
	location := entities.Location{
		X:     x,
		Y:     feetY,
		Z:     z,
		Yaw:   lastLocation.Yaw,
		Pitch: lastLocation.Pitch,
	}
	player.SetLocation(&location)
	player.SetOnGround(onGround)
}

func RunPlayTick(*entities_manager.EntityManager) {
	ProcessPacketQueue()
}
