package serverbound

import (
	player "HexaUtils/entities/player"
	"HexaUtils/packets"
	packet_utils "HexaUtils/packets/utils"
)

type ActionID int

const (
	StartSneaking         ActionID = 0
	StopSneaking          ActionID = 1
	LeaveBed              ActionID = 2
	StartSprinting        ActionID = 3
	StopSprinting         ActionID = 4
	StartJumpWithHorse    ActionID = 5
	StopJumpWithHorse     ActionID = 6
	OpenVehicleInventory  ActionID = 7
	StartFlyingWithElytra ActionID = 8
)

type PlayerCommandPacket_1_21 struct {
	PacketID          int
	ServerBoundPacket bool
	ProtocolVersion   int
	State             player.ClientState
	EntityID          int32
	ActionID          ActionID
	JumpBoost         int32
}

func (p *PlayerCommandPacket_1_21) GetEntityID() int32 {
	return p.EntityID
}

func (p *PlayerCommandPacket_1_21) SetEntityID(entityID int32) {
	p.EntityID = entityID
}

func (p *PlayerCommandPacket_1_21) GetActionID() ActionID {
	return p.ActionID
}

func (p *PlayerCommandPacket_1_21) SetActionID(actionID ActionID) {
	p.ActionID = actionID
}

func (p *PlayerCommandPacket_1_21) GetJumpBoost() int32 {
	return p.JumpBoost
}

func (p *PlayerCommandPacket_1_21) SetJumpBoost(jumpBoost int32) {
	p.JumpBoost = jumpBoost
}

func (p *PlayerCommandPacket_1_21) GetProtocolVersion() int {
	return p.ProtocolVersion
}

func (p *PlayerCommandPacket_1_21) GetPacketID() int {
	return p.PacketID
}

func (p *PlayerCommandPacket_1_21) GetState() player.ClientState {
	return p.State
}

func (p *PlayerCommandPacket_1_21) IsServerBound() bool {
	return p.ServerBoundPacket
}

func NewPlayerCommandPacket_1_21(entityID int32, actionID ActionID, jumpBoost int32) PlayerCommandPacket_1_21 {
	return PlayerCommandPacket_1_21{
		PacketID:          0x25,
		ServerBoundPacket: true,
		ProtocolVersion:   767,
		State:             player.Play,
		EntityID:          entityID,
		ActionID:          actionID,
		JumpBoost:         jumpBoost,
	}
}

func ReadPlayerCommandPacket_1_21(packet *packet_utils.PacketReader) (PlayerCommandPacket_1_21, bool) {
	entityID, err := packet.ReadVarInt()
	if err != nil {
		return PlayerCommandPacket_1_21{}, false
	}
	actionID, err := packet.ReadVarInt()
	if err != nil {
		return PlayerCommandPacket_1_21{}, false
	}
	jumpBoost, err := packet.ReadVarInt()
	if err != nil {
		return PlayerCommandPacket_1_21{}, false
	}

	return PlayerCommandPacket_1_21{
		PacketID:          0x25,
		ServerBoundPacket: true,
		ProtocolVersion:   767,
		State:             player.Play,
		EntityID:          entityID,
		ActionID:          ActionID(actionID),
		JumpBoost:         jumpBoost,
	}, true
}

func (p PlayerCommandPacket_1_21) GetPacket(player player.Player) *packets.Packet {
	//packet := packet_utils.NewPacketWriter()
	packet := player.GetPacketWritter()
	packet.Reset()
	packet.WriteVarInt(p.EntityID)
	packet.WriteVarInt(int32(p.ActionID))
	packet.WriteVarInt(p.JumpBoost)

	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"PlayerCommandPacket",
		packet.GetPacketBuffer(),
		p.ServerBoundPacket,
		p.State)
	return real_packet
}
