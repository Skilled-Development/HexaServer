package packets

import (
	"HexaProtocol_1_21/packets/clientbound"
	serverbound "HexaProtocol_1_21/packets/serverbound"
	"HexaUtils/entities/player"
	"HexaUtils/packets"
	packet_utils "HexaUtils/packets/utils"
	config "HexaUtils/server/config"
	"HexaUtils/server/data"
	debugger "HexaUtils/utils"
	"fmt"
	"time"
)

func ReadConfigurationStatePacket(server_config config.ServerConfig, p player.Player, length int32, packet_id int32, pack *packets.PacketReader) {
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
		clientbound_plugin_message_packet.GetPacket().Send(p)
		p.AddAlreadySendPacket("ClientboundPluginMessage")
	}
	if !p.ContainsAlreadySendPacket("ClientboundKnownPacks") {
		kwonwPack := packet_utils.NewKnownPack("minecraft", "core", "1.0.0")
		kwonwPacks := []packet_utils.KnownPack{*kwonwPack}
		clientbound_known_packs_packet := clientbound.NewClientboundKnownPacks_1_21(kwonwPacks)
		clientbound_known_packs_packet.GetPacket().Send(p)
		p.AddAlreadySendPacket("ClientboundKnownPacks")
	}

}

func handle_serverbound_acknowledge_finish_configuration_packet(p player.Player, pack *packets.PacketReader) {
	p.SetClientState(player.Play)
	packet := clientbound.NewPlayPacket_1_21(p)
	packet.GetPacket().Send(p)
	flags := clientbound.CreateFlags(clientbound.FlagPitch, clientbound.FlagYaw, clientbound.FlagX, clientbound.FlagY, clientbound.FlagZ)
	synchronize_player_position_packet := clientbound.NewSynchronizePlayerPositionPacketFromPlayer_1_21(p, flags)
	synchronize_player_position_packet.GetPacket().Send(p)

	game_event_packet := clientbound.NewGameEventPacket_1_21(clientbound.StartWaitingForLevelChunksEvent, 0)
	game_event_packet.GetPacket().Send(p)

	set_center_chunk_packet := clientbound.NewSetCenterChunkPacket_1_21(0, 1)
	set_center_chunk_packet.GetPacket().Send(p)

	startTimeSendChunks := time.Now()

	debugger.SetDebugTest(false)
	for x := -32; x < 32; x++ {
		for z := -32; z < 32; z++ {
			region := data.GetRegionsLoadedList()[0]
			chunk, _ := region.ReadChunk(x, z)
			if chunk == nil {
				continue
			}
			/*if len(chunk.BlockEntities) > 0 {
				debugger.SetDebugTest(true)
			} else {
				debugger.SetDebugTest(false)
			}*/
			debugger.PrintForDebug("--------------------------------------------------------------------------")
			debugger.PrintForDebug("Current chunk x:", x, "z:", z)
			chunkPacket := clientbound.NewChunkDataAndUpdateLightPacket_1_21_FromChunkStruct(chunk, p)
			chunkPacket.GetPacket().Send(p)
		}
	}
	debugger.SetDebugTest(false)
	println("Time to send chunks:", time.Since(startTimeSendChunks).Milliseconds(), "ms")

}

func handle_serverbound_known_packs_packet(p player.Player, pack *packets.PacketReader) {
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
		registryPacket.GetPacket().Send(p)
	}
	packet := packets.NewPacketWriter()
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

func handle_client_information_packet(p player.Player, pack *packets.PacketReader) {
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

func handle_serverbound_plugin_message_configuration_packet(pack *packets.PacketReader) {
	plugin_message_configuration_packet := serverbound.ReadPluginMessageConfigurationPacket_1_21(pack)
	if plugin_message_configuration_packet == nil {
		fmt.Println("Error reading PluginMessageConfigurationPacket")
		return
	}

}
