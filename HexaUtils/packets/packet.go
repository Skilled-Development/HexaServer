package packets

import (
	"HexaUtils/entities/player"
	"fmt"
	"time"
)

type Packet struct {
	PacketID          int
	ProtocolVersion   int
	Name              string
	ServerBoundPacket bool
	PacketBuffer      []byte
	State             player.ClientState
}

func (p *Packet) GetClientState() player.ClientState {
	return p.State
}

func (p *Packet) GetPacketID() int {
	return p.PacketID
}

func (p *Packet) GetProtocolVersion() int {
	return p.ProtocolVersion
}

func (p *Packet) GetName() string {
	return p.Name
}

func (p *Packet) IsServerBound() bool {
	return p.ServerBoundPacket
}

func (p *Packet) GetPacketBuffer() []byte {
	return p.PacketBuffer
}

func NewPacket(packet_id int, protocol_version int, name string, packetbuffer []byte, serverbound bool, state player.ClientState) *Packet {
	return &Packet{
		PacketID:          packet_id,
		ProtocolVersion:   protocol_version,
		Name:              name,
		State:             state,
		PacketBuffer:      packetbuffer,
		ServerBoundPacket: serverbound,
	}
}

func (p *Packet) GetLogMessage() string {
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	buffer := p.GetPacketBuffer()
	direction := "S -> C"
	if p.IsServerBound() {
		direction = "C -> S"
	}
	buffer_as_hexadecimal := string(buffer)

	packet_log := "[" + currentTime + "] [" + direction + "] (" + p.GetName() + " | " + fmt.Sprintf("%d", p.GetProtocolVersion()) + ") " + buffer_as_hexadecimal

	return packet_log
}

func (p *Packet) Build() []byte {
	buffer := p.GetPacketBuffer()
	packet_length := len(buffer)
	otherWriter := NewPacketWriter()
	otherWriter.WriteVarInt(int32(packet_length))
	otherWriter.buffer = append(otherWriter.buffer, buffer...)
	return otherWriter.buffer
}

func (p *Packet) Send(pl player.Player) {
	// Obtener la conexión del jugador
	conn := pl.GetConn()

	newBuffer := p.Build()

	// Escribir el paquete en la conexión
	(*conn).Write(newBuffer) //<-- Error here

	if p.GetClientState() == player.Play {
		pl.AddPacketLogToPlayer(p.GetLogMessage())
	}
}

