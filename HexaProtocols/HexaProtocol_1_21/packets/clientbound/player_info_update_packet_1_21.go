package clientbound

import (
	"HexaUtils/entities/player"
	"HexaUtils/packets"

	"github.com/google/uuid"
)

type PlayerInfoUpdatePacket_1_21 struct {
	PacketID          int
	ServerBoundPacket bool
	ProtocolVersion   int
	State             player.ClientState
	Actions           PlayerInfoActions
	NumberOfPlayers   int32
	Players           []PlayerInfoEntry
}

type PlayerInfoActions byte

const (
	AddPlayerAction         PlayerInfoActions = 0x01
	InitializeChatAction    PlayerInfoActions = 0x02
	UpdateGameModeAction    PlayerInfoActions = 0x04
	UpdateListedAction      PlayerInfoActions = 0x08
	UpdateLatencyAction     PlayerInfoActions = 0x10
	UpdateDisplayNameAction PlayerInfoActions = 0x20
)

type PlayerInfoEntry struct {
	UUID          uuid.UUID
	PlayerActions []PlayerActionData
}

type PlayerActionData struct {
	ActionType            PlayerInfoActions
	AddPlayerData         AddPlayerData
	InitializeChatData    InitializeChatData
	UpdateGameModeData    UpdateGameModeData
	UpdateListedData      UpdateListedData
	UpdateLatencyData     UpdateLatencyData
	UpdateDisplayNameData UpdateDisplayNameData
}

type AddPlayerData struct {
	Name       string
	Properties []PlayerProperty
}

type PlayerProperty struct {
	Name      string
	Value     string
	IsSigned  bool
	Signature string
}

type InitializeChatData struct {
	HasSignatureData       bool
	PublickeyExpiryTime    int64
	EncodedPublicKeySize   int32
	EncodedPublicKey       []byte
	PublickeySignatureSize int32
	PublickeySignature     []byte
}

type UpdateGameModeData struct {
	GameMode int32
}
type UpdateListedData struct {
	Listed bool
}
type UpdateLatencyData struct {
	Ping int32
}
type UpdateDisplayNameData struct {
	HasDisplayName bool
	DisplayName    *string
}

func NewPlayerInfoUpdatePacket_1_21(actions PlayerInfoActions, players []PlayerInfoEntry) *PlayerInfoUpdatePacket_1_21 {
	return &PlayerInfoUpdatePacket_1_21{
		PacketID:          0x3E,
		ServerBoundPacket: false,
		ProtocolVersion:   767,
		State:             player.Play,
		Actions:           actions,
		NumberOfPlayers:   int32(len(players)),
		Players:           players,
	}
}

func (p *PlayerInfoUpdatePacket_1_21) GetActions() PlayerInfoActions {
	return p.Actions
}

func (p *PlayerInfoUpdatePacket_1_21) SetActions(actions PlayerInfoActions) {
	p.Actions = actions
}

func (p *PlayerInfoUpdatePacket_1_21) GetNumberOfPlayers() int32 {
	return p.NumberOfPlayers
}

func (p *PlayerInfoUpdatePacket_1_21) SetNumberOfPlayers(numberOfPlayers int32) {
	p.NumberOfPlayers = numberOfPlayers
}

func (p *PlayerInfoUpdatePacket_1_21) GetPlayers() []PlayerInfoEntry {
	return p.Players
}

func (p *PlayerInfoUpdatePacket_1_21) SetPlayers(players []PlayerInfoEntry) {
	p.Players = players
}

func (p *PlayerInfoUpdatePacket_1_21) GetPacket() *packets.Packet {
	packet := packets.NewPacketWriter()
	packet.WriteVarInt(int32(p.PacketID))
	packet.WriteByte(byte(p.Actions))
	packet.WriteVarInt(p.NumberOfPlayers)
	for _, player := range p.Players {
		packet.WriteUUID(player.UUID)
		for _, action := range player.PlayerActions {
			switch action.ActionType {
			case AddPlayerAction:
				packet.WriteString(action.AddPlayerData.Name)
				packet.WriteVarInt(int32(len(action.AddPlayerData.Properties)))
				for _, prop := range action.AddPlayerData.Properties {
					packet.WriteString(prop.Name)
					packet.WriteString(prop.Value)
					packet.WriteBoolean(prop.IsSigned)
					if prop.IsSigned {
						packet.WriteString(prop.Signature)
					}
				}
			case InitializeChatAction:
				packet.WriteBoolean(action.InitializeChatData.HasSignatureData)
				if action.InitializeChatData.HasSignatureData {
					packet.WriteLong(action.InitializeChatData.PublickeyExpiryTime)
					packet.WriteVarInt(action.InitializeChatData.EncodedPublicKeySize)
					packet.WriteByteArray(action.InitializeChatData.EncodedPublicKey)
					packet.WriteVarInt(action.InitializeChatData.PublickeySignatureSize)
					packet.WriteByteArray(action.InitializeChatData.PublickeySignature)
				}

			case UpdateGameModeAction:
				packet.WriteVarInt(action.UpdateGameModeData.GameMode)

			case UpdateListedAction:
				packet.WriteBoolean(action.UpdateListedData.Listed)

			case UpdateLatencyAction:
				packet.WriteVarInt(action.UpdateLatencyData.Ping)

			case UpdateDisplayNameAction:
				packet.WriteBoolean(action.UpdateDisplayNameData.HasDisplayName)
				if action.UpdateDisplayNameData.HasDisplayName && action.UpdateDisplayNameData.DisplayName != nil {
					packet.WriteString(*action.UpdateDisplayNameData.DisplayName)
				}
			}
		}

	}

	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"PlayerInfoUpdatePacket",
		packet.GetPacketBuffer(),
		p.ServerBoundPacket,
		p.State)
	return real_packet
}
