package serverbound

import (
	player "HexaUtils/entities/player"
	"HexaUtils/packets"
	packet_utils "HexaUtils/packets/utils"
)

type ServerboundPlayerPositionPacket_1_21 struct {
	PacketID          int
	ServerBoundPacket bool
	ProtocolVersion   int
	State             player.ClientState
	X                 float64
	FeetY             float64
	Z                 float64
	OnGround          bool
}

func (p *ServerboundPlayerPositionPacket_1_21) GetX() float64 {
	return p.X
}

func (p *ServerboundPlayerPositionPacket_1_21) SetX(x float64) {
	p.X = x
}

func (p *ServerboundPlayerPositionPacket_1_21) GetFeetY() float64 {
	return p.FeetY
}

func (p *ServerboundPlayerPositionPacket_1_21) SetFeetY(feetY float64) {
	p.FeetY = feetY
}

func (p *ServerboundPlayerPositionPacket_1_21) GetZ() float64 {
	return p.Z
}

func (p *ServerboundPlayerPositionPacket_1_21) SetZ(z float64) {
	p.Z = z
}

func (p *ServerboundPlayerPositionPacket_1_21) GetOnGround() bool {
	return p.OnGround
}

func (p *ServerboundPlayerPositionPacket_1_21) SetOnGround(onGround bool) {
	p.OnGround = onGround
}

func (p *ServerboundPlayerPositionPacket_1_21) GetProtocolVersion() int {
	return p.ProtocolVersion
}

func (p *ServerboundPlayerPositionPacket_1_21) GetPacketID() int {
	return p.PacketID
}

func (p *ServerboundPlayerPositionPacket_1_21) GetState() player.ClientState {
	return p.State
}

func (p *ServerboundPlayerPositionPacket_1_21) IsServerBound() bool {
	return p.ServerBoundPacket
}

func NewServerboundPlayerPositionPacket_1_21(x float64, feetY float64, z float64, onGround bool) ServerboundPlayerPositionPacket_1_21 {
	return ServerboundPlayerPositionPacket_1_21{
		PacketID:          0x1A, // Packet ID for Serverbound Player Position
		ServerBoundPacket: false,
		ProtocolVersion:   767, // or your specific protocol version
		State:             player.Play,
		X:                 x,
		FeetY:             feetY,
		Z:                 z,
		OnGround:          onGround,
	}
}

func ReadServerboundPlayerPositionPacket_1_21(packet packet_utils.PacketReader) (ServerboundPlayerPositionPacket_1_21, bool) {
	x, err := packet.ReadDouble()
	if err != nil {
		return ServerboundPlayerPositionPacket_1_21{}, false
	}
	feetY, err := packet.ReadDouble()
	if err != nil {
		return ServerboundPlayerPositionPacket_1_21{}, false
	}
	z, err := packet.ReadDouble()
	if err != nil {
		return ServerboundPlayerPositionPacket_1_21{}, false
	}
	onGround, err := packet.ReadBoolean()
	if err != nil {
		return ServerboundPlayerPositionPacket_1_21{}, false
	}
	return ServerboundPlayerPositionPacket_1_21{
		PacketID:          0x1A,
		ServerBoundPacket: true,
		ProtocolVersion:   767,
		State:             player.Play,
		X:                 x,
		FeetY:             feetY,
		Z:                 z,
		OnGround:          onGround,
	}, true
}

func (p ServerboundPlayerPositionPacket_1_21) GetPacket(player player.Player) *packets.Packet {
	//packet := packet_utils.NewPacketWriter()
	packet := player.GetPacketWritter()
	packet.Reset()
	packet.WriteVarInt(int32(p.PacketID))
	packet.WriteDouble(p.X)
	packet.WriteDouble(p.FeetY)
	packet.WriteDouble(p.Z)
	packet.WriteBoolean(p.OnGround)

	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"ServerboundPlayerPositionPacket_1_21",
		packet.GetPacketBuffer(),
		p.ServerBoundPacket,
		p.State)
	return real_packet
}
