// serverbound/interact.go
package serverbound

import (
	entities "HexaUtils/entities"
	player "HexaUtils/entities/player"
	player_package "HexaUtils/entities/player"
	"HexaUtils/packets"
	packet_utils "HexaUtils/packets/utils"
)

type InteractPacket_1_21 struct {
	PacketID          int
	ServerBoundPacket bool
	ProtocolVersion   int
	State             player.ClientState
	EntityID          int32
	Type              entities.InteractionType
	TargetX           *float32
	TargetY           *float32
	TargetZ           *float32
	Hand              *player_package.Hand
	Sneaking          bool
}

func (p *InteractPacket_1_21) GetEntityID() int32 {
	return p.EntityID
}

func (p *InteractPacket_1_21) SetEntityID(entityID int32) {
	p.EntityID = entityID
}

func (p *InteractPacket_1_21) GetType() entities.InteractionType {
	return p.Type
}

func (p *InteractPacket_1_21) SetType(interactionType entities.InteractionType) {
	p.Type = interactionType
}

func (p *InteractPacket_1_21) GetTargetX() *float32 {
	return p.TargetX
}

func (p *InteractPacket_1_21) SetTargetX(targetX *float32) {
	p.TargetX = targetX
}

func (p *InteractPacket_1_21) GetTargetY() *float32 {
	return p.TargetY
}

func (p *InteractPacket_1_21) SetTargetY(targetY *float32) {
	p.TargetY = targetY
}

func (p *InteractPacket_1_21) GetTargetZ() *float32 {
	return p.TargetZ
}

func (p *InteractPacket_1_21) SetTargetZ(targetZ *float32) {
	p.TargetZ = targetZ
}

func (p *InteractPacket_1_21) GetHand() *player_package.Hand {
	return p.Hand
}

func (p *InteractPacket_1_21) SetHand(hand *player_package.Hand) {
	p.Hand = hand
}

func (p *InteractPacket_1_21) GetSneaking() bool {
	return p.Sneaking
}

func (p *InteractPacket_1_21) SetSneaking(sneaking bool) {
	p.Sneaking = sneaking
}

func (p *InteractPacket_1_21) GetProtocolVersion() int {
	return p.ProtocolVersion
}

func (p *InteractPacket_1_21) GetPacketID() int {
	return p.PacketID
}

func (p *InteractPacket_1_21) GetState() player.ClientState {
	return p.State
}
func (p *InteractPacket_1_21) IsServerBound() bool {
	return p.ServerBoundPacket
}
func NewInteractPacket_1_21(entityID int32, interactionType entities.InteractionType, targetX *float32, targetY *float32, targetZ *float32, hand *player_package.Hand, sneaking bool) InteractPacket_1_21 {
	return InteractPacket_1_21{
		PacketID:          0x16,
		ServerBoundPacket: true,
		ProtocolVersion:   767,
		State:             player.Play,
		EntityID:          entityID,
		Type:              interactionType,
		TargetX:           targetX,
		TargetY:           targetY,
		TargetZ:           targetZ,
		Hand:              hand,
		Sneaking:          sneaking,
	}
}

func ReadInteractPacket_1_21(packet packet_utils.PacketReader) (InteractPacket_1_21, bool) {
	entityID, err := packet.ReadVarInt()
	if err != nil {
		return InteractPacket_1_21{}, false
	}

	interactionType, err := packet.ReadVarInt()
	if err != nil {
		return InteractPacket_1_21{}, false
	}

	var targetX *float32
	var targetY *float32
	var targetZ *float32

	if entities.InteractionType(interactionType) == entities.InteractAt {
		x, err := packet.ReadFloat()
		if err != nil {
			return InteractPacket_1_21{}, false
		}
		y, err := packet.ReadFloat()
		if err != nil {
			return InteractPacket_1_21{}, false
		}
		z, err := packet.ReadFloat()
		if err != nil {
			return InteractPacket_1_21{}, false
		}
		targetX = &x
		targetY = &y
		targetZ = &z
	}

	var hand *player_package.Hand
	if entities.InteractionType(interactionType) == entities.Interact || entities.InteractionType(interactionType) == entities.InteractAt {
		handVarInt, err := packet.ReadVarInt()
		if err != nil {
			return InteractPacket_1_21{}, false
		}
		h := player_package.Hand(handVarInt)
		hand = &h
	}

	sneaking, err := packet.ReadBoolean()
	if err != nil {
		return InteractPacket_1_21{}, false
	}

	return InteractPacket_1_21{
		PacketID:          0x16,
		ServerBoundPacket: true,
		ProtocolVersion:   767,
		State:             player.Play,
		EntityID:          entityID,
		Type:              entities.InteractionType(interactionType),
		TargetX:           targetX,
		TargetY:           targetY,
		TargetZ:           targetZ,
		Hand:              hand,
		Sneaking:          sneaking,
	}, true
}

func (p InteractPacket_1_21) GetPacket(player player.Player) *packets.Packet {
	//packet := packets.NewPacketWriter()
	packet := player.GetPacketWritter()
	packet.Reset()
	packet.WriteVarInt(p.EntityID)
	packet.WriteVarInt(int32(p.Type))

	if p.Type == entities.InteractAt {
		packet.WriteFloat(*p.TargetX)
		packet.WriteFloat(*p.TargetY)
		packet.WriteFloat(*p.TargetZ)
	}

	if p.Type == entities.Interact || p.Type == entities.InteractAt {
		packet.WriteVarInt(int32(*p.Hand))
	}
	packet.WriteBoolean(p.Sneaking)

	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"InteractPacket",
		packet.GetPacketBuffer(),
		p.ServerBoundPacket,
		p.State)
	return real_packet
}
