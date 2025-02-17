package packets

import (
	"HexaUtils/entities/player"
	"HexaUtils/packets/utils"
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
	packetLength := len(buffer)

	otherWriter := utils.NewPacketWriter()
	otherWriter.WriteVarInt(int32(packetLength))
	otherBuffer := otherWriter.GetPacketBuffer()
	finalBuffer := make([]byte, 0, len(otherBuffer)+len(buffer))
	finalBuffer = append(finalBuffer, otherBuffer...)
	finalBuffer = append(finalBuffer, buffer...)
	return finalBuffer
}
func (p *Packet) Send(pl player.Player) {
	// Obtener la conexión del jugador
	conn := pl.GetConn()
	newBuffer := p.Build()
	(*conn).Write(newBuffer)
}
