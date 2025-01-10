package packets

import (
	current_entities "HexaProtocol_1_21/entities"
	"HexaProtocol_1_21/packets/clientbound"
	clientbound_play "HexaProtocol_1_21/packets/clientbound/play"
	serverbound "HexaProtocol_1_21/packets/serverbound"
	entities_manager "HexaServer/entities/manager"
	"HexaUtils/entities/player"
	"HexaUtils/packets"
	packet_utils "HexaUtils/packets/utils"
	chunk_generator "HexaUtils/regionreader/generator"
	config "HexaUtils/server/config"
	debugger "HexaUtils/utils"
	"fmt"
	"time"
)

func ReadConfigurationStatePacket(server_config config.ServerConfig, p player.Player, length int32, packet_id int32, pack *packet_utils.PacketReader) {
	switch packet_id {
	case 0x00:
		handle_client_information_packet(p, pack)
	case 0x02:
		handle_serverbound_plugin_message_configuration_packet(pack)
	case 0x03:
		handle_serverbound_acknowledge_finish_configuration_packet(p, pack)
	case 0x07:
		handle_serverbound_known_packs_packet(p, pack)
	default:
		fmt.Println("Unknown packet ID in Configuration state: ", packet_id)
	}
	if !p.ContainsAlreadySendPacket("ClientboundPluginMessage") {
		clientbound_plugin_message_packet := clientbound.NewClientboundPluginMessagePacket_1_21("minecraft:brand", []byte("HexaServer"))
		clientbound_plugin_message_packet.GetPacket(p).Send(p)
		p.AddAlreadySendPacket("ClientboundPluginMessage")
	}
	if !p.ContainsAlreadySendPacket("ClientboundKnownPacks") {
		kwonwPack := packet_utils.NewKnownPack("minecraft", "core", "1.0.0")
		kwonwPacks := []packet_utils.KnownPack{*kwonwPack}
		clientbound_known_packs_packet := clientbound.NewClientboundKnownPacks_1_21(kwonwPacks)
		clientbound_known_packs_packet.GetPacket(p).Send(p)
		p.AddAlreadySendPacket("ClientboundKnownPacks")
	}

}

func handle_serverbound_acknowledge_finish_configuration_packet(p player.Player, pack *packet_utils.PacketReader) {
	//TODO: PLAYER JOIN WORLD EVENT
	p.SetClientState(player.Play)
	packet := clientbound.NewPlayPacket_1_21(p)
	packet.GetPacket(p).Send(p)
	synchronize_player_position_packet := clientbound.NewSynchronizePlayerPositionPacketFromPlayer_1_21(p, 0)
	synchronize_player_position_packet.GetPacket(p).Send(p)

	game_event_packet := clientbound.NewGameEventPacket_1_21(clientbound.StartWaitingForLevelChunksEvent, 0)
	game_event_packet.GetPacket(p).Send(p)

	set_center_chunk_packet := clientbound.NewSetCenterChunkPacket_1_21(0, 1)
	set_center_chunk_packet.GetPacket(p).Send(p)

	startTimeSendChunks := time.Now()

	debugger.SetDebugTest(false)
	//region := data.GetRegionsLoadedList()[0]
	//loadedChunk, _ := region.ReadChunk(0, 0)

	/*for x := -32; x < 32; x++ {
		for z := -32; z < 32; z++ {
			region := data.GetRegionsLoadedList()[0]
			chunk, _ := region.ReadChunk(x, z)
			if chunk == nil {
				continue
			}
			//if len(chunk.BlockEntities) > 0 {
				//debugger.SetDebugTest(true)
			//} else {
			//	debugger.SetDebugTest(false)
			//}
			debugger.PrintForDebug("--------------------------------------------------------------------------")
			debugger.PrintForDebug("Current chunk x:", x, "z:", z)
			chunkPacket := clientbound.NewChunkDataAndUpdateLightPacket_1_21_FromChunkStruct(chunk, p)
			chunkPacket.GetPacket(p).Send(p)
		}
	}
	*/

	zero_zero_chunk := chunk_generator.GenerateChunk(0, 0, 0)
	chunkPacket := clientbound.NewChunkDataAndUpdateLightPacket_1_21_FromChunkStruct(zero_zero_chunk, p)
	chunkPacket.GetPacket().Send(p)

	/*

		debugger.SetDebugTest(false)*/
	for x := -10; x < 10; x++ {
		for y := -10; y < 10; y++ {
			/*file, err := os.OpenFile("debug.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println("Error opening file:", err)
				return
			}
			defer file.Close()
			writer := bufio.NewWriter(file)
			fmt.Fprintln(writer, "GENERATING CHUNK X:", x, "Y:", y)*/
			if x == 0 && y == 0 {
				continue
			}
			chunk := chunk_generator.GenerateChunk(int32(x), int32(y), 0)
			/*if x == 0 && y == 0 {
				fmt.Fprintln(writer, "----------------")
				fmt.Fprintln(writer, "LOADED CHUNK X POS:", loadedChunk.XPos, "Z POS:", loadedChunk.ZPos, " Y POS:", loadedChunk.YPos)
				fmt.Fprintln(writer, "GENERATED CHUNK X POS:", chunk.XPos, "Z POS:", chunk.ZPos, " Y POS:", chunk.YPos)
				fmt.Fprintln(writer, "Example of heigthmaps, for loaded chunk")
				for key, value := range loadedChunk.Heightmaps {
					fmt.Fprintln(writer, " - Key:", key, "Value:", value)
				}
				fmt.Fprintln(writer, "Example of heigthmaps, for generated chunk")
				for key, value := range chunk.Heightmaps {
					fmt.Fprintln(writer, " - Key:", key, "Value:", value)
				}
				section := 5
				fmt.Fprintln(writer, "Example of section", section, "palette for loaded chunk")
				for key, value := range loadedChunk.Sections[section].BlockStates.Palette {
					fmt.Fprintln(writer, " - Key:", key, "Value:", value)
					fmt.Fprintln(writer, "  - BlockState Name:", value.Name)
					for propKey, propValue := range value.Properties {
						fmt.Fprintln(writer, "    - Property:", propKey, "Value:", propValue)
					}
				}
				fmt.Fprintln(writer, "Example of section", section, "data for loaded chunk")
				for key, value := range loadedChunk.Sections[section].BlockStates.Data {
					fmt.Fprintln(writer, " - Key:", key, "Value:", value)

				}
				fmt.Fprintln(writer, "Example of section", section, "palette for generated chunk")
				for key, value := range chunk.Sections[section].BlockStates.Palette {
					fmt.Fprintln(writer, " - Key:", key, "Value:", value)
				}
				fmt.Fprintln(writer, "Example of section", section, "data for generated chunk")
				for key, value := range chunk.Sections[section].BlockStates.Data {
					fmt.Fprintln(writer, " - Key:", key, "Value:", value)
				}
			}
			writer.Flush()*/
			chunkPacket := clientbound.NewChunkDataAndUpdateLightPacket_1_21_FromChunkStruct(chunk, p)
			chunkPacket.GetPacket().Send(p)
		}
	}
	println("Time to send chunks:", time.Since(startTimeSendChunks).Milliseconds(), "ms")

	addPlayerPacket := clientbound.NewPlayerInfoUpdatePacket_1_21(clientbound.AddPlayerAction, []clientbound.PlayerInfoEntry{getPlayerInfoEntry(p)})
	others_players := entities_manager.EntityManagerInstance.GetPlayersExcept(p.GetEntityId())
	for _, other_player := range others_players {
		sendMessage("Player "+p.GetName()+" joined the game", other_player)
		sendMessage("Tu nombre es "+p.GetName(), p)

		pLocation := p.GetLocation()
		spawn_player_packet := clientbound_play.NewSpawnEntityPacket_1_21(
			int32(p.GetEntityId()),
			p.GetUUID(),
			current_entities.PLAYER,
			pLocation.X,
			pLocation.Y,
			pLocation.Z,
			pLocation.Pitch,
			pLocation.Yaw,
			pLocation.Yaw,
			0,
			0,
			0,
			0,
		)
		addPlayerPacket.GetPacket(other_player).Send(other_player)
		spawn_player_packet.GetPacket(other_player).Send(other_player)
		other_player.RemoveSeeingEntity(p.GetEntityId())
		other_player.AddSeeingEntity(p.GetEntityId())

		addOtherPlayerPacket := clientbound.NewPlayerInfoUpdatePacket_1_21(clientbound.AddPlayerAction, []clientbound.PlayerInfoEntry{getPlayerInfoEntry(other_player)})
		other_location := other_player.GetLocation()
		spawn_other_palyer := clientbound_play.NewSpawnEntityPacket_1_21(
			int32(other_player.GetEntityId()),
			other_player.GetUUID(),
			current_entities.PLAYER,
			other_location.X,
			other_location.Y,
			other_location.Z,
			other_location.Pitch,
			other_location.Yaw,
			other_location.Yaw,
			0,
			0,
			0,
			0,
		)
		addOtherPlayerPacket.GetPacket(p).Send(p)
		spawn_other_palyer.GetPacket(p).Send(p)
		p.RemoveSeeingEntity(other_player.GetEntityId())
		p.AddSeeingEntity(other_player.GetEntityId())
		//sendMessage("Player "+p.GetName()+" spawned at "+fmt.Sprintf("%.2f", pLocation.X)+", "+fmt.Sprintf("%.2f", pLocation.Y)+", "+fmt.Sprintf("%.2f", pLocation.Z), other_player)
	}
	//sendMessage("Bienvenido a HexaServer "+p.GetName()+", Disfruta!!", p)

	go func() {
		time.Sleep(1 * time.Second)
		others_players := entities_manager.EntityManagerInstance.GetPlayers()
		for _, other_player := range others_players {
			teleportPacket := clientbound_play.NewTeleportEntityPacket_1_21(
				int32(other_player.GetEntityId()),
				other_player.GetLocation().X,
				other_player.GetLocation().Y,
				other_player.GetLocation().Z,
				byte(other_player.GetLocation().Yaw),
				byte(other_player.GetLocation().Pitch),
				other_player.IsOnGround(),
			)
			teleportPacket.GetPacket(p).Send(p)

			teleportPacket2 := clientbound_play.NewTeleportEntityPacket_1_21(
				int32(p.GetEntityId()),
				p.GetLocation().X,
				p.GetLocation().Y,
				p.GetLocation().Z,
				byte(p.GetLocation().Yaw),
				byte(p.GetLocation().Pitch),
				p.IsOnGround(),
			)
			teleportPacket2.GetPacket(other_player).Send(other_player)
		}

	}()

}

func sendMessage(s string, p player.Player) {
	packet := clientbound.NewSystemChatMessagePacket_1_21(s, false)
	packet.GetPacket(p).Send(p)
}

func getPlayerInfoEntry(p player.Player) clientbound.PlayerInfoEntry {
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
		AddPlayerData: add_player_data_action,
	}
	player_info_entry := clientbound.PlayerInfoEntry{
		UUID:          p.GetUUID(),
		PlayerActions: []clientbound.PlayerActionData{player_action},
	}
	return player_info_entry
}

func handle_serverbound_known_packs_packet(p player.Player, pack *packet_utils.PacketReader) {
	serverbound_known_packs_packet := serverbound.ReadServerboundKnownPacks_1_21(pack)
	if serverbound_known_packs_packet == nil {
		fmt.Println("Error reading ServerboundKnownPacksPacket")
		return
	}
	knownpacks := serverbound_known_packs_packet.GetKnownPacks()
	for _, knownpack := range knownpacks {
		//TODO: Implement KnownPack
		fmt.Println("KnownPack: ", knownpack.GetNamespace(), knownpack.GetID(), knownpack.GetVersion())
	}
	registryManager := config.RegistriesManagerInstance
	for _, registry := range registryManager.Registries {
		registryPacket := clientbound.NewRegistryDataPacket_1_21(registry)
		registryPacket.GetPacket(p).Send(p)
	}
	packet := packet_utils.NewPacketWriter()
	packet.WriteVarInt(int32(0x03))
	real_packet := packets.NewPacket(0x03,
		p.GetProtocolVersion(),
		"ClientInformationPacket",
		packet.GetPacketBuffer(),
		true,
		player.Configuration,
	)
	real_packet.Send(p)
}

func handle_client_information_packet(p player.Player, pack *packet_utils.PacketReader) {
	client_information_packet := serverbound.ReadClientInformationPacket_1_21(pack)
	if client_information_packet == nil {
		fmt.Println("Error reading ClientInformationPacket")
		return
	}
	p.SetLocale(client_information_packet.GetLocale())
	p.SetViewDistance(client_information_packet.GetViewDistance())
	p.SetChatMode(client_information_packet.GetChatMode())
	p.SetChatColors(client_information_packet.ChatColors)
	p.SetDisplayedSkinParts(client_information_packet.DisplayedSkinParts)
	p.SetMainHand(client_information_packet.MainHand)
	p.SetEnableTextFilter(client_information_packet.EnableTextFilter)
	p.SetAllowServerListings(client_information_packet.AllowServerListings)
}

func handle_serverbound_plugin_message_configuration_packet(pack *packet_utils.PacketReader) {
	plugin_message_configuration_packet := serverbound.ReadPluginMessageConfigurationPacket_1_21(pack)
	if plugin_message_configuration_packet == nil {
		fmt.Println("Error reading PluginMessageConfigurationPacket")
		return
	}

}
