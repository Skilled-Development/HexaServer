package packet

import (
	protocol_1_21 "HexaProtocol_1_21/packets"
	"HexaServer/entities/player"
	utils_player "HexaUtils/entities/player"
	hexapackets "HexaUtils/packets"
	config "HexaUtils/server/config"
	"fmt"
	"log"
)

// PacketReader es un lector de paquetes que usa un buffer.
type PlayerPacketReader struct {
	packet_reader_1_21 *protocol_1_21.PlayerPacketReader_1_21
}

// NewPacketReader crea un nuevo lector de paquetes.
func NewPlayerPacketReader() *PlayerPacketReader {
	return &PlayerPacketReader{
		packet_reader_1_21: protocol_1_21.NewPlayerPacketReader_1_21(),
	}
}

func (r *PlayerPacketReader) ReadPlayerPacket(packetReader *hexapackets.PacketReader, player *player.Player, server_config *config.ServerConfig) {
	playerProtocolVersion := player.GetProtocolVersion()
	if playerProtocolVersion == 0 {
		r.readHandshakePacket(player, packetReader)
	} else {
		switch playerProtocolVersion {
		case 767:
			r.packet_reader_1_21.ReadPacket(player, packetReader, server_config) // Usar la instancia que tienes
		}
	}
}
func (r *PlayerPacketReader) readHandshakePacket(p *player.Player, packetReader *hexapackets.PacketReader) {
	conn := p.Conn
	length, err := packetReader.ReadVarInt()
	if err != nil {
		log.Printf("Error al leer Length del cliente %s: %v\n", (*conn).RemoteAddr(), err)
		return
	}
	packet_id, err := packetReader.ReadVarInt()
	if err != nil {
		log.Printf("Error al leer Packet ID del cliente %s: %v\n", (*conn).RemoteAddr(), err)
		return
	}
	fmt.Println("Length: ", length, " Packet ID: ", packet_id)
	protocol_version, err := packetReader.ReadVarInt()
	if err != nil {
		log.Printf("Error al leer Protocol Version del cliente %s: %v\n", (*conn).RemoteAddr(), err)
		return
	}
	fmt.Println("Protocol Version: ", protocol_version)
	// Leer el nombre del jugador
	server_address, err := packetReader.ReadString()
	if err != nil {
		log.Printf("Error al leer Server Address del cliente %s: %v\n", (*conn).RemoteAddr(), err)
		return
	}
	fmt.Println("Server Address: ", server_address)
	server_port, err := packetReader.ReadUnsignedShort()
	if err != nil {
		log.Printf("Error al leer Server Port del cliente %s: %v\n", (*conn).RemoteAddr(), err)
		return
	}
	fmt.Println("Server Port: ", server_port)
	next_state, err := packetReader.ReadVarInt()
	if err != nil {
		log.Printf("Error al leer Next State del cliente %s: %v\n", (*conn).RemoteAddr(), err)
		return
	}
	fmt.Println("Next State: ", next_state)
	// Set state to status

	p.SetProtocolVersion(int(protocol_version)) // Establecer la versión del protocolo
	// Aquí cambiamos el estado del jugador a 'Status'
	if next_state == 1 { // El valor 1 corresponde a 'Status' según la definición en el archivo client_state.go
		p.SetClientState(utils_player.Status) // Cambiar el estado a "Status"
		fmt.Println("Estado del jugador cambiado a Status.")
	} else if next_state == 2 {
		p.SetClientState(utils_player.Login) // Cambiar el estado a "Login"
		fmt.Println("Estado del jugador cambiado a Login.")
	}
}
