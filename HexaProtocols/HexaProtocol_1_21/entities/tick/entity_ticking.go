package tick

/*clientbound_play "HexaProtocol_1_21/packets/clientbound/play"
entities_manager "HexaServer/entities/manager"*/

func TickEntities(tick int) {
	/*if tick%20 == 0 {
		others_players := entities_manager.EntityManagerInstance.GetPlayers()
		for _, other_player := range others_players {
			for _, p := range others_players {
				teleportPacket := clientbound_play.NewTeleportEntityPacket_1_21(
					int32(other_player.GetEntityId()),
					other_player.GetLocation().X,
					other_player.GetLocation().Y,
					other_player.GetLocation().Z,
					byte(other_player.GetLocation().Yaw),
					byte(other_player.GetLocation().Pitch),
					other_player.IsOnGround(),
				)
				teleportPacket.GetPacket().Send(p)

				teleportPacket2 := clientbound_play.NewTeleportEntityPacket_1_21(
					int32(p.GetEntityId()),
					p.GetLocation().X,
					p.GetLocation().Y,
					p.GetLocation().Z,
					byte(p.GetLocation().Yaw),
					byte(p.GetLocation().Pitch),
					p.IsOnGround(),
				)
				teleportPacket2.GetPacket().Send(other_player)
			}
		}
	}*/
}
