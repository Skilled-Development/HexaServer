package clientbound

import (
	"HexaUtils/entities/player"
	"HexaUtils/packets"
	"math/rand"
	"time"
)

type SyncPositionFlag byte

const (
	FlagX     SyncPositionFlag = 0x01
	FlagY     SyncPositionFlag = 0x02
	FlagZ     SyncPositionFlag = 0x04
	FlagPitch SyncPositionFlag = 0x08
	FlagYaw   SyncPositionFlag = 0x10
)

// Function to create a bitmask from multiple flags
func CreateFlags(flags ...SyncPositionFlag) byte {
	var result byte
	for _, flag := range flags {
		result |= byte(flag)
	}
	return result
}
func GenerateRandomTeleportID() int32 {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	return rand.Int31()
}

type SynchronizePlayerPositionPacket_1_21 struct {
	PacketID          int
	ServerBoundPacket bool
	ProtocolVersion   int
	State             player.ClientState
	X                 float64
	Y                 float64
	Z                 float64
	Yaw               float32
	Pitch             float32
	Flags             byte
	TeleportID        int32
}

func (p *SynchronizePlayerPositionPacket_1_21) GetProtocolVersion() int {
	return p.ProtocolVersion
}

func (p *SynchronizePlayerPositionPacket_1_21) GetPacketID() int {
	return p.PacketID
}

func (p *SynchronizePlayerPositionPacket_1_21) GetState() player.ClientState {
	return p.State
}

func (p *SynchronizePlayerPositionPacket_1_21) IsServerBound() bool {
	return p.ServerBoundPacket
}

func NewSynchronizePlayerPositionPacket_1_21(x float64, y float64, z float64, yaw float32, pitch float32, flags byte) *SynchronizePlayerPositionPacket_1_21 {
	return &SynchronizePlayerPositionPacket_1_21{
		PacketID:          0x40,
		ServerBoundPacket: false,
		ProtocolVersion:   767, // Assuming 1.21 protocol version
		State:             player.Play,
		X:                 x,
		Y:                 y,
		Z:                 z,
		Yaw:               yaw,
		Pitch:             pitch,
		Flags:             flags,
		TeleportID:        GenerateRandomTeleportID(),
	}
}

func (p *SynchronizePlayerPositionPacket_1_21) GetTeleportId() int32 {
	return p.TeleportID
}

func NewSynchronizePlayerPositionPacketFromPlayer_1_21(p player.Player, flags byte) *SynchronizePlayerPositionPacket_1_21 {
	randomTpid := GenerateRandomTeleportID()
	p.SetTeleportID(randomTpid)
	return &SynchronizePlayerPositionPacket_1_21{
		PacketID:          0x40,
		ServerBoundPacket: false,
		ProtocolVersion:   767, // Assuming 1.21 protocol version
		State:             player.Play,
		X:                 p.GetLocation().GetX(),
		Y:                 p.GetLocation().GetY(),
		Z:                 p.GetLocation().GetZ(),
		Yaw:               float32(p.GetLocation().GetYaw()),
		Pitch:             float32(p.GetLocation().GetPitch()),
		Flags:             flags,
		TeleportID:        randomTpid,
	}
}

func (p *SynchronizePlayerPositionPacket_1_21) GetPacket() *packets.Packet {
	packet := packets.NewPacketWriter()
	packet.WriteVarInt(int32(p.PacketID))

	packet.WriteDouble(p.X)
	packet.WriteDouble(p.Y)
	packet.WriteDouble(p.Z)
	packet.WriteFloat(p.Yaw)
	packet.WriteFloat(p.Pitch)
	packet.WriteByte(p.Flags)
	packet.WriteVarInt(p.TeleportID)

	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"SynchronizePlayerPositionPacket",
		packet.GetPacketBuffer(),
		p.ServerBoundPacket,
		p.State)
	return real_packet
}
