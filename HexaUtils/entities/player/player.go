package player

import (
	utils_entities "HexaUtils/entities"
	"net"

	"github.com/google/uuid"
)

type Player interface {
	GetEntityId() int64
	GetName() string
	SetName(name string)
	GetHealth() float64
	GetUUID() uuid.UUID
	SetUUID(uuid uuid.UUID)
	GetClientState() ClientState
	SetClientState(state ClientState)
	GetProtocolVersion() int
	GetConn() *net.Conn
	AddPacketLogToPlayer(log string)
	SetLocale(locale string)
	SetViewDistance(viewDistance byte)
	SetChatMode(chatMode int32)
	SetChatColors(chatColors bool)
	SetDisplayedSkinParts(displayedSkinParts byte)
	SetMainHand(mainHand int32)
	SetEnableTextFilter(enableTextFilter bool)
	SetAllowServerListings(allowServerListings bool)
	SetAlreadySendPackets(packets []string)
	GetAlreadySendPackets() []string
	AddAlreadySendPacket(packet string)
	ContainsAlreadySendPacket(packet string) bool
	GetGamemode() GameMode
	GetLocation() *utils_entities.Location
	SetLocation(location *utils_entities.Location)
	SetTeleportID(id int32)
	GetTeleportID() int32
	IsOnGround() bool
	SetOnGround(onGround bool)
	GetEntityType() utils_entities.EntityType
	GetWorldID() int64
	GetSkinValue() string
	GetSkinSignature() string
	NeedsKeepAlivePacket() bool
	GenerateKeepAliveID() int64
	GetKeepAliveID() int64
	GetSeeingEntityList() []int64
	SetSeeingEntityList(entities []int64)
	AddSeeingEntity(entityID int64)
	RemoveSeeingEntity(entityID int64)
	IsSeeingEntity(entityID int64) bool
	SetJumpBoost(jumpBoost int32)
	GetJumpBoost() int32
	SetSneaking(sneaking bool)
	IsSneaking() bool
	SetSprinting(sprinting bool)
	IsSprinting() bool
}
