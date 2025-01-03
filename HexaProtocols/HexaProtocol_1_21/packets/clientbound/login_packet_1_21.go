package clientbound

import (
	"HexaUtils/entities/player"
	"HexaUtils/packets"
	"HexaUtils/registries"
	"HexaUtils/server/config"
	"crypto/sha256"
	"encoding/binary"
	"math/rand"
	"time"
)

type PlayPacket_1_21 struct {
	PacketID          int
	ServerBoundPacket bool
	ProtocolVersion   int
	State             player.ClientState
	EntityId          int64
	Hardcore          bool
	GameMode          player.GameMode
}

func (p *PlayPacket_1_21) GetEntityId() int64 {
	return p.EntityId
}

func (p *PlayPacket_1_21) GetProtocolVersion() int {
	return p.ProtocolVersion
}

func (p *PlayPacket_1_21) GetPacketID() int {
	return p.PacketID
}

func (p *PlayPacket_1_21) GetState() player.ClientState {
	return p.State
}

func (p *PlayPacket_1_21) IsServerBound() bool {
	return p.ServerBoundPacket
}

func (p *PlayPacket_1_21) GetGameMode() player.GameMode {
	return p.GameMode
}

func (p *PlayPacket_1_21) IsHardcore() bool {
	return p.Hardcore
}

func (p *PlayPacket_1_21) SetEntityId(entityId int64) {
	p.EntityId = entityId
}

func (p *PlayPacket_1_21) SetGameMode(gameMode player.GameMode) {
	p.GameMode = gameMode
}

func NewPlayPacket_1_21(p player.Player) *PlayPacket_1_21 {
	return &PlayPacket_1_21{
		PacketID:          0x2B,
		ServerBoundPacket: false,
		ProtocolVersion:   767,
		State:             player.Login,
		EntityId:          p.GetEntityId(),
		Hardcore:          p.GetGamemode() == player.Hardcore,
		GameMode:          p.GetGamemode(),
	}
}

func generateHashedSeed() int64 {
	// 1. Generar un seed aleatorio (en este caso, un string)
	rand.Seed(time.Now().UnixNano()) // Inicializar el generador de números aleatorios
	seedBytes := make([]byte, 32)
	rand.Read(seedBytes)
	seedString := string(seedBytes)

	// 2. Calcular el hash SHA-256 del seed
	hasher := sha256.New()
	hasher.Write([]byte(seedString))
	hashBytes := hasher.Sum(nil)

	// 3. Tomar los primeros 8 bytes del hash
	first8Bytes := hashBytes[:8]

	// 4. Convertir los 8 bytes a un int64 (Long en Java)
	hashedSeed := int64(binary.LittleEndian.Uint64(first8Bytes))

	return hashedSeed
}

func (p *PlayPacket_1_21) GetPacket() *packets.Packet {
	packet := packets.NewPacketWriter()
	packet.WriteVarInt(int32(p.PacketID))   // packet id
	packet.WriteInt(int32(p.GetEntityId())) // entity id
	packet.WriteBoolean(p.IsHardcore())     // hardcore

	dimensionsCount := len(registries.DimensionTypeRegistryInstance.GetEntries())
	packet.WriteVarInt(int32(dimensionsCount)) // dimension count
	for _, dimension := range registries.DimensionTypeRegistryInstance.GetEntries() {
		packet.WriteIdentifier(dimension.GetName()) // dimension id
	}

	packet.WriteVarInt(int32(config.ServerConfigInstance.GetMaxPlayers()))         // max players
	packet.WriteVarInt(int32(config.ServerConfigInstance.GetViewDistance()))       // view distance
	packet.WriteVarInt(int32(config.ServerConfigInstance.GetSimulationDistance())) // simulation distance
	packet.WriteBoolean(false)                                                     // reduced debug info
	packet.WriteBoolean(true)                                                      // enable respawn screen
	packet.WriteBoolean(false)                                                     // do limited crafting

	packet.WriteVarInt(0)
	packet.WriteIdentifier("minecraft:overworld") // dimension name
	packet.WriteLong(2041223745)                  // hashed seed
	packet.WriteUnsignedByte(1)                   // gamemode
	packet.WriteByte(byte(0))                     // previous gamemode
	packet.WriteBoolean(false)                    // debug world
	packet.WriteBoolean(false)                    // flat world
	packet.WriteBoolean(false)                    // has death location
	packet.WriteVarInt(0)                         // portal cooldown
	packet.WriteBoolean(true)                     // enforces secure chat

	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"LoginPacket",
		packet.GetPacketBuffer(),
		p.ServerBoundPacket,
		p.State)
	return real_packet
}
