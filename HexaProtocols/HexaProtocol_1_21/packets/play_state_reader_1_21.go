package packets

import (
	"HexaProtocol_1_21/packets/clientbound"
	clientbound_play "HexaProtocol_1_21/packets/clientbound/play"
	"HexaProtocol_1_21/packets/serverbound"
	serverbound_play "HexaProtocol_1_21/packets/serverbound/play"
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
	case serverbound_play.SetPlayerRotationPacket_1_21:
		handle_serverbound_player_rotation_packet(player, p)
	case serverbound_play.PlayerCommandPacket_1_21:
		handle_serverbound_player_command_packet(player, p)
	case serverbound_play.SwingArmPacket_1_21:
		handle_serverbound_swing_arm_packet(player, p)
	default:
		fmt.Println("Paquete desconocido recibido", packet.GetPacket().GetPacketID())
		fmt.Printf("Packet type: %T\n", packet)
	}
}

func handle_serverbound_swing_arm_packet(playerP player.Player, p serverbound_play.SwingArmPacket_1_21) {
	hand := p.GetHand()
	allPlayers := entities_manager.EntityManagerInstance.GetPlayersExcept(playerP.GetEntityId())
	animation := clientbound_play.SwingMainArm
	if hand == player.OffHand {
		animation = clientbound_play.SwingOffhand
	}
	entity_animation_packet := clientbound_play.NewEntityAnimationPacket_1_21(int32(playerP.GetEntityId()), animation)
	for _, other := range allPlayers {
		if other.GetClientState().String() == "Play" && other.IsSeeingEntity(playerP.GetEntityId()) {
			entity_animation_packet.GetPacket().Send(other)
		}
	}
}

func ReadPlayStatePacket(server_config *config.ServerConfig, p player.Player, length int32, packet_id int32, pack packets.PacketReader, reader *PlayerPacketReader_1_21) {
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
	case 0x1C:
		create_player_rotation_packet(p, pack)
	case 0x25:
		create_player_command_packet(p, pack)
	case 0x36:
		create_swing_arm_packet(p, pack)
	default:
		println("Unknown packet ID:", fmt.Sprintf("0x%X", packet_id))
		println("Unknown packet ID:", packet_id)
	}

}

func create_swing_arm_packet(p player.Player, pack packets.PacketReader) {
	swing_arm_packet, ok := serverbound_play.ReadSwingArmPacket_1_21(pack)
	if !ok {
		return
	}
	EnqueuePacket(p, swing_arm_packet)
}

func create_player_command_packet(p player.Player, pack packets.PacketReader) {
	player_command_packet, ok := serverbound_play.ReadPlayerCommandPacket_1_21(&pack)
	if !ok {
		return
	}
	EnqueuePacket(p, player_command_packet)
}

func handle_serverbound_player_command_packet(player player.Player, p serverbound_play.PlayerCommandPacket_1_21) {
	actionID := p.GetActionID()
	jumpBoost := p.GetJumpBoost()
	player.SetJumpBoost(jumpBoost)
	switch actionID {
	case serverbound_play.StartSneaking:
		player.SetSneaking(true)
		metadata := []*clientbound_play.MetadataEntry{
			{
				Index: 6,
				Type:  clientbound_play.MetadataTypePose,
				Value: int32(5),
			},
			{
				Index: 0xff,
			},
		}
		metadataPacket := clientbound_play.NewSetEntityMetadataPacket_1_21(
			int32(player.GetEntityId()),
			metadata,
		)
		allPlayers := entities_manager.EntityManagerInstance.GetPlayersExcept(player.GetEntityId())
		for _, other := range allPlayers {
			if other.GetClientState().String() == "Play" && other.IsSeeingEntity(player.GetEntityId()) {
				metadataPacket.GetPacket().Send(other)
			}
		}
	case serverbound_play.StopSneaking:
		player.SetSneaking(false)
		metadata := []*clientbound_play.MetadataEntry{
			{
				Index: 6,
				Type:  clientbound_play.MetadataTypePose,
				Value: int32(0),
			},
			{
				Index: 0xff,
			},
		}
		metadataPacket := clientbound_play.NewSetEntityMetadataPacket_1_21(
			int32(player.GetEntityId()),
			metadata,
		)
		allPlayers := entities_manager.EntityManagerInstance.GetPlayersExcept(player.GetEntityId())
		for _, other := range allPlayers {
			if other.GetClientState().String() == "Play" && other.IsSeeingEntity(player.GetEntityId()) {
				metadataPacket.GetPacket().Send(other)
			}
		}
	case serverbound_play.StartSprinting:
		player.SetSprinting(true)
	case serverbound_play.StopSprinting:
		player.SetSprinting(false)
		//TODO: add more
	}
}

func create_player_rotation_packet(p player.Player, pack packets.PacketReader) {
	player_rotation_packet, ok := serverbound_play.ReadSetPlayerRotationPacket_1_21(&pack)
	if !ok {
		return
	}
	EnqueuePacket(p, player_rotation_packet)
}

func handle_serverbound_player_rotation_packet(player player.Player, p serverbound_play.SetPlayerRotationPacket_1_21) {
	yaw := p.GetYaw()
	pitch := p.GetPitch()
	onGround := p.GetOnGround()
	lastLocation := player.GetLocation()
	location := entities.Location{
		X:     lastLocation.X,
		Y:     lastLocation.Y,
		Z:     lastLocation.Z,
		Yaw:   yaw,
		Pitch: pitch,
	}
	player.SetLocation(&location)
	player.SetOnGround(onGround)
	updateHeadRotationPacket := clientbound_play.NewSetHeadRotationPacket_1_21(int32(player.GetEntityId()), byte(yaw*256/360))
	updateEntityRotionPacket := clientbound_play.NewUpdateEntityRotationPacket_1_21(int32(player.GetEntityId()), byte(yaw*256/360), byte(pitch*256/360), onGround)
	allPlayers := entities_manager.EntityManagerInstance.GetPlayersExcept(player.GetEntityId())
	for _, other := range allPlayers {
		if other.GetClientState().String() == "Play" && other.IsSeeingEntity(player.GetEntityId()) {
			updateEntityRotionPacket.GetPacket().Send(other)
			updateHeadRotationPacket.GetPacket().Send(other)
		}
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
	lastX := last_location.X
	lastY := last_location.Y
	lastZ := last_location.Z
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
	p.SetLocation(&location)
	p.SetOnGround(onGround)

	last_chunk_x := int32(last_location.X / 16)
	last_chunk_z := int32(last_location.Z / 16)
	current_chunk_x := int32(location.X / 16)
	current_chunk_z := int32(location.Z / 16)
	if last_chunk_x != current_chunk_x || last_chunk_z != current_chunk_z {
		set_center_chunk_packet := clientbound.NewSetCenterChunkPacket_1_21(current_chunk_x, current_chunk_z)
		set_center_chunk_packet.GetPacket().Send(p)
	}

	deltaX := x*4096 - lastX*4096
	deltaY := feetY*4096 - lastY*4096
	deltaZ := z*4096 - lastZ*4096
	updatePositionPacket := clientbound_play.NewUpdatePositionAndRotationPacket_1_21(
		int32(p.GetEntityId()),
		int16(deltaX),
		int16(deltaY),
		int16(deltaZ),
		byte(yaw*256/360),
		byte(pitch*256/360),
		p.IsOnGround(),
	)
	updateHeadRotationPacket := clientbound_play.NewSetHeadRotationPacket_1_21(int32(p.GetEntityId()), byte(yaw*256/360))
	allPlayers := entities_manager.EntityManagerInstance.GetPlayersExcept(p.GetEntityId())
	for _, other := range allPlayers {
		if other.GetClientState().String() == "Play" && other.IsSeeingEntity(p.GetEntityId()) {
			updatePositionPacket.GetPacket().Send(other)
			updateHeadRotationPacket.GetPacket().Send(other)
		}
	}
}

func create_serverbound_keep_alive_packet(p player.Player, pack packets.PacketReader) {
	serverbound_keep_alive_packet := serverbound.ReadKeepAlivePlayPacket_1_21(pack)
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
	lastX := lastLocation.X
	lastY := lastLocation.Y
	lastZ := lastLocation.Z
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
	deltaX := x*4096 - lastX*4096
	deltaY := feetY*4096 - lastY*4096
	deltaZ := z*4096 - lastZ*4096
	updatePositionPacket := clientbound_play.NewUpdatePositionPacket_1_21(int32(player.GetEntityId()), int16(deltaX), int16(deltaY), int16(deltaZ), player.IsOnGround())
	allPlayers := entities_manager.EntityManagerInstance.GetPlayersExcept(player.GetEntityId())
	for _, p := range allPlayers {
		if p.GetClientState().String() == "Play" && p.IsSeeingEntity(player.GetEntityId()) {
			updatePositionPacket.GetPacket().Send(p)
		}
	}

}

func RunPlayTick(*entities_manager.EntityManager) {
	ProcessPacketQueue()
}
