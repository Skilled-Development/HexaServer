// serverbound/playerrotation.go
package serverbound

import (
	player "HexaUtils/entities/player"
	"HexaUtils/packets"
	packet_utils "HexaUtils/packets/utils"
)

type SetPlayerRotationPacket_1_21 struct {
	PacketID          int
	ServerBoundPacket bool
	ProtocolVersion   int
	State             player.ClientState
	Yaw               float32
	Pitch             float32
	OnGround          bool
}

func (p *SetPlayerRotationPacket_1_21) GetYaw() float32 {
	return p.Yaw
}

func (p *SetPlayerRotationPacket_1_21) SetYaw(yaw float32) {
	p.Yaw = yaw
}

func (p *SetPlayerRotationPacket_1_21) GetPitch() float32 {
	return p.Pitch
}

func (p *SetPlayerRotationPacket_1_21) SetPitch(pitch float32) {
	p.Pitch = pitch
}

func (p *SetPlayerRotationPacket_1_21) GetOnGround() bool {
	return p.OnGround
}

func (p *SetPlayerRotationPacket_1_21) SetOnGround(onGround bool) {
	p.OnGround = onGround
}

func (p *SetPlayerRotationPacket_1_21) GetProtocolVersion() int {
	return p.ProtocolVersion
}

func (p *SetPlayerRotationPacket_1_21) GetPacketID() int {
	return p.PacketID
}

func (p *SetPlayerRotationPacket_1_21) GetState() player.ClientState {
	return p.State
}

func (p *SetPlayerRotationPacket_1_21) IsServerBound() bool {
	return p.ServerBoundPacket
}

func NewSetPlayerRotationPacket_1_21(yaw float32, pitch float32, onGround bool) SetPlayerRotationPacket_1_21 {
	return SetPlayerRotationPacket_1_21{
		PacketID:          0x1C,
		ServerBoundPacket: true,
		ProtocolVersion:   767,
		State:             player.Play,
		Yaw:               yaw,
		Pitch:             pitch,
		OnGround:          onGround,
	}
}

func ReadSetPlayerRotationPacket_1_21(packet *packet_utils.PacketReader) (SetPlayerRotationPacket_1_21, bool) {
	yaw, err := packet.ReadFloat()
	if err != nil {
		return SetPlayerRotationPacket_1_21{}, false
	}
	pitch, err := packet.ReadFloat()
	if err != nil {
		return SetPlayerRotationPacket_1_21{}, false
	}
	onGround, err := packet.ReadBoolean()
	if err != nil {
		return SetPlayerRotationPacket_1_21{}, false
	}
	return SetPlayerRotationPacket_1_21{
		PacketID:          0x1C,
		ServerBoundPacket: true,
		ProtocolVersion:   767,
		State:             player.Play,
		Yaw:               yaw,
		Pitch:             pitch,
		OnGround:          onGround,
	}, true
}

func (p SetPlayerRotationPacket_1_21) GetPacket(player player.Player) *packets.Packet {
	//packet := packet_utils.NewPacketWriter()
	packet := player.GetPacketWritter()
	packet.Reset()
	packet.WriteVarInt(int32(p.PacketID))
	packet.WriteFloat(p.Yaw)
	packet.WriteFloat(p.Pitch)
	packet.WriteBoolean(p.OnGround)

	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"SetPlayerRotationPacket",
		packet.GetPacketBuffer(),
		p.ServerBoundPacket,
		p.State)
	return real_packet
}
