package serverbound

import (
	"HexaUtils/entities/player"
	"HexaUtils/packets"
)

type ConfirmTeleportation_1_21 struct {
	PacketID          int
	ServerBoundPacket bool
	ProtocolVersion   int
	State             player.ClientState
	TeleportID        int32
}

func (p ConfirmTeleportation_1_21) GetProtocolVersion() int {
	return p.ProtocolVersion
}

func (p ConfirmTeleportation_1_21) GetPacketID() int {
	return p.PacketID
}

func (p ConfirmTeleportation_1_21) GetState() player.ClientState {
	return p.State
}

func (p ConfirmTeleportation_1_21) IsServerBound() bool {
	return p.ServerBoundPacket
}

func NewConfirmTeleportation_1_21(teleportID int32) ConfirmTeleportation_1_21 {
	return ConfirmTeleportation_1_21{
		PacketID:          0x00,
		ServerBoundPacket: true,
		ProtocolVersion:   767,
		State:             player.Play,
		TeleportID:        teleportID,
	}
}

func ReadConfirmTeleportation_1_21(packet packets.PacketReader) ConfirmTeleportation_1_21 {
	teleportID, err := packet.ReadVarInt()
	if err != nil {
		return ConfirmTeleportation_1_21{}
	}

	return ConfirmTeleportation_1_21{
		PacketID:          0x00,
		ServerBoundPacket: true,
		ProtocolVersion:   767,
		State:             player.Play,
		TeleportID:        teleportID,
	}
}

func (p ConfirmTeleportation_1_21) GetTeleportId() int32 {
	return p.TeleportID
}

func (p ConfirmTeleportation_1_21) GetPacket() *packets.Packet {
	packet := packets.NewPacketWriter()
	packet.WriteVarInt(int32(p.PacketID))
	packet.WriteVarInt(p.TeleportID)

	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"ConfirmTeleportation",
		packet.GetPacketBuffer(),
		p.ServerBoundPacket,
		p.State)
	return real_packet
}
