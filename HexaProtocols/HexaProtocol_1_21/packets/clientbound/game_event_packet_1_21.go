package clientbound

import (
	"HexaUtils/entities/player"
	"HexaUtils/packets"
)

type GameEventPacket_1_21 struct {
	PacketID          int
	ServerBoundPacket bool
	ProtocolVersion   int
	State             player.ClientState
	Event             GameEventType
	Value             float32
}

type GameEventType byte

const (
	NoRespawnBlockAvailableEvent        GameEventType = 0
	BeginRainingEvent                   GameEventType = 1
	EndRainingEvent                     GameEventType = 2
	ChangeGameModeEvent                 GameEventType = 3
	WinGameEvent                        GameEventType = 4
	DemoEvent                           GameEventType = 5
	ArrowHitPlayerEvent                 GameEventType = 6
	RainLevelChangeEvent                GameEventType = 7
	ThunderLevelChangeEvent             GameEventType = 8
	PlayPufferfishStingSoundEvent       GameEventType = 9
	PlayElderGuardianMobAppearanceEvent GameEventType = 10
	EnableRespawnScreenEvent            GameEventType = 11
	LimitedCraftingEvent                GameEventType = 12
	StartWaitingForLevelChunksEvent     GameEventType = 13
)

// Specific value types for some events
type GameMode int32

const (
	Survival  GameMode = 0
	Creative  GameMode = 1
	Adventure GameMode = 2
	Spectator GameMode = 3
)

type WinGame int32

const (
	JustRespawnPlayer           WinGame = 0
	RollCreditsAndRespawnPlayer WinGame = 1
)

type DemoEventValue int32

const (
	ShowWelcomeToDemoScreen DemoEventValue = 0
	TellMovementControls    DemoEventValue = 101
	TellJumpControl         DemoEventValue = 102
	TellInventoryControl    DemoEventValue = 103
	TellThatTheDemoIsOver   DemoEventValue = 104
)

type EnableRespawnScreen int32

const (
	EnableRespawnScreenValue EnableRespawnScreen = 0
	ImmediatelyRespawnValue  EnableRespawnScreen = 1
)

type LimitedCrafting int32

const (
	DisableLimitedCrafting LimitedCrafting = 0
	EnableLimitedCrafting  LimitedCrafting = 1
)

func NewGameEventPacket_1_21(event GameEventType, value float32) *GameEventPacket_1_21 {
	return &GameEventPacket_1_21{
		PacketID:          0x22,
		ServerBoundPacket: false,
		ProtocolVersion:   767,
		State:             player.Play,
		Event:             event,
		Value:             value,
	}
}

func (p *GameEventPacket_1_21) GetEvent() GameEventType {
	return p.Event
}

func (p *GameEventPacket_1_21) SetEvent(event GameEventType) {
	p.Event = event
}

func (p *GameEventPacket_1_21) GetValue() float32 {
	return p.Value
}

func (p *GameEventPacket_1_21) SetValue(value float32) {
	p.Value = value
}

func (p *GameEventPacket_1_21) GetPacket() *packets.Packet {
	packet := packets.NewPacketWriter()
	packet.WriteVarInt(int32(p.PacketID))
	packet.WriteByte(byte(p.Event))
	packet.WriteFloat(p.Value)
	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"GameEventPacket",
		packet.GetPacketBuffer(),
		p.ServerBoundPacket,
		p.State)
	return real_packet
}
