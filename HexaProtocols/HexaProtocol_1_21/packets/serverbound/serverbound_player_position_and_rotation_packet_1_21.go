package serverbound

import (
	player "HexaUtils/entities/player"
	"HexaUtils/packets"
	packet_utils "HexaUtils/packets/utils"
)

type ServerboundPlayerPositionAndRotation_1_21 struct {
	PacketID          int
	ServerBoundPacket bool
	ProtocolVersion   int
	State             player.ClientState
	X                 float64
	FeetY             float64
	Z                 float64
	Yaw               float32
	Pitch             float32
	OnGround          bool
}

func (p ServerboundPlayerPositionAndRotation_1_21) GetX() float64 {
	return p.X
}

func (p ServerboundPlayerPositionAndRotation_1_21) SetX(x float64) {
	p.X = x
}

func (p ServerboundPlayerPositionAndRotation_1_21) GetFeetY() float64 {
	return p.FeetY
}

func (p ServerboundPlayerPositionAndRotation_1_21) SetFeetY(feetY float64) {
	p.FeetY = feetY
}

func (p ServerboundPlayerPositionAndRotation_1_21) GetZ() float64 {
	return p.Z
}

func (p ServerboundPlayerPositionAndRotation_1_21) SetZ(z float64) {
	p.Z = z
}

func (p ServerboundPlayerPositionAndRotation_1_21) GetYaw() float32 {
	return p.Yaw
}

func (p ServerboundPlayerPositionAndRotation_1_21) SetYaw(yaw float32) {
	p.Yaw = yaw
}

func (p ServerboundPlayerPositionAndRotation_1_21) GetPitch() float32 {
	return p.Pitch
}

func (p ServerboundPlayerPositionAndRotation_1_21) SetPitch(pitch float32) {
	p.Pitch = pitch
}

func (p ServerboundPlayerPositionAndRotation_1_21) GetOnGround() bool {
	return p.OnGround
}

func (p ServerboundPlayerPositionAndRotation_1_21) SetOnGround(onGround bool) {
	p.OnGround = onGround
}

func (p ServerboundPlayerPositionAndRotation_1_21) GetProtocolVersion() int {
	return p.ProtocolVersion
}

func (p ServerboundPlayerPositionAndRotation_1_21) GetPacketID() int {
	return p.PacketID
}

func (p ServerboundPlayerPositionAndRotation_1_21) GetState() player.ClientState {
	return p.State
}

func (p ServerboundPlayerPositionAndRotation_1_21) IsServerBound() bool {
	return p.ServerBoundPacket
}

func NewServerboundPlayerPositionAndRotation_1_21(x float64, feetY float64, z float64, yaw float32, pitch float32, onGround bool) ServerboundPlayerPositionAndRotation_1_21 {
	return ServerboundPlayerPositionAndRotation_1_21{
		PacketID:          0x1B,
		ServerBoundPacket: false,
		ProtocolVersion:   767,
		State:             player.Play,
		X:                 x,
		FeetY:             feetY,
		Z:                 z,
		Yaw:               yaw,
		Pitch:             pitch,
		OnGround:          onGround,
	}
}

func ReadServerboundPlayerPositionAndRotation_1_21(packet packet_utils.PacketReader) (ServerboundPlayerPositionAndRotation_1_21, bool) {
	x, err := packet.ReadDouble()
	if err != nil {
		return ServerboundPlayerPositionAndRotation_1_21{}, false
	}
	feetY, err := packet.ReadDouble()
	if err != nil {
		return ServerboundPlayerPositionAndRotation_1_21{}, false
	}
	z, err := packet.ReadDouble()
	if err != nil {
		return ServerboundPlayerPositionAndRotation_1_21{}, false
	}
	yaw, err := packet.ReadFloat()
	if err != nil {
		return ServerboundPlayerPositionAndRotation_1_21{}, false
	}
	pitch, err := packet.ReadFloat()
	if err != nil {
		return ServerboundPlayerPositionAndRotation_1_21{}, false
	}
	onGround, err := packet.ReadBoolean()
	if err != nil {
		return ServerboundPlayerPositionAndRotation_1_21{}, false
	}
	return ServerboundPlayerPositionAndRotation_1_21{
		PacketID:          0x1B,
		ServerBoundPacket: true,
		ProtocolVersion:   767,
		State:             player.Play,
		X:                 x,
		FeetY:             feetY,
		Z:                 z,
		Yaw:               yaw,
		Pitch:             pitch,
		OnGround:          onGround,
	}, true
}

func (p ServerboundPlayerPositionAndRotation_1_21) GetPacket(player player.Player) *packets.Packet {
	//packet := packet_utils.NewPacketWriter()
	packet := player.GetPacketWritter()
	packet.Reset()
	packet.WriteVarInt(int32(p.PacketID))
	packet.WriteDouble(p.X)
	packet.WriteDouble(p.FeetY)
	packet.WriteDouble(p.Z)
	packet.WriteFloat(p.Yaw)
	packet.WriteFloat(p.Pitch)
	packet.WriteBoolean(p.OnGround)

	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"ServerboundPlayerPositionAndRotation_1_21",
		packet.GetPacketBuffer(),
		p.ServerBoundPacket,
		p.State)
	return real_packet
}
