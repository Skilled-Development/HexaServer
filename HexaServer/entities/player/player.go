package player

import (
	"HexaUtils/entities"
	"HexaUtils/entities/player"
	"bufio"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Player struct {
	EntityId            int64
	Name                string
	Health              float64
	UUID                uuid.UUID
	PlayerState         player.ClientState
	Conn                *net.Conn
	ProtocolVersion     int
	mu                  sync.Mutex
	Locale              string
	ViewDistance        byte
	ChatMode            int32
	ChatColors          bool
	DisplayedSkinParts  byte
	MainHand            int32
	EnableTextFilter    bool
	AllowServerListings bool
	AlReadySendPackets  []string
	Gamemode            player.GameMode
	Location            *entities.Location
	TeleportID          int32
	OnGround            bool
	WorldID             int64
	DimensionID         int32
	SkinValue           string
	SkinSignature       string
	LastKeepAliveTime   int64
	SeeingEntityList    []int64
}

func (p *Player) GetSeeingEntityList() []int64 {
	return p.SeeingEntityList
}

func (p *Player) SetSeeingEntityList(list []int64) {
	p.SeeingEntityList = list
}

func (p *Player) AddSeeingEntity(entityID int64) {
	p.SeeingEntityList = append(p.SeeingEntityList, entityID)
}

func (p *Player) IsSeeingEntity(entityID int64) bool {
	for _, id := range p.SeeingEntityList {
		if id == entityID {
			return true
		}
	}
	return false
}

func (p *Player) RemoveSeeingEntity(entityID int64) {
	for i, id := range p.SeeingEntityList {
		if id == entityID {
			p.SeeingEntityList = append(p.SeeingEntityList[:i], p.SeeingEntityList[i+1:]...)
			break
		}
	}
}

func (p *Player) NeedsKeepAlivePacket() bool {
	const keepAliveInterval = 15 * 1000 // 15 seconds in milliseconds

	currentTime := time.Now().UnixMilli()
	return currentTime-p.LastKeepAliveTime >= keepAliveInterval
}

func (p *Player) GetKeepAliveID() int64 {
	return p.LastKeepAliveTime
}

func (p *Player) GenerateKeepAliveID() int64 {
	p.LastKeepAliveTime = time.Now().UnixMilli()
	return p.LastKeepAliveTime
}

func (p *Player) GetSkinValue() string {
	return p.SkinValue
}

func (p *Player) GetSkinSignature() string {
	return p.SkinSignature
}

func (p *Player) GetDimensionID() int32 {
	return p.DimensionID
}

func (p *Player) GetWorldID() int64 {
	return p.WorldID
}

func (p *Player) GetEntityType() entities.EntityType {
	return entities.Player
}

func (p *Player) SetOnGround(onGround bool) {
	p.OnGround = onGround
}

func (p *Player) IsOnGround() bool {
	return p.OnGround
}

func (p *Player) SetTeleportID(id int32) {
	p.TeleportID = id
}

func (p *Player) GetTeleportID() int32 {
	return p.TeleportID
}

func (p *Player) GetLocation() *entities.Location {
	return p.Location
}

func (p *Player) SetLocation(location *entities.Location) {
	p.Location = location
}

func (p *Player) GetGamemode() player.GameMode {
	return p.Gamemode
}

func (p *Player) GetEntityId() int64 {
	return p.EntityId
}

func (p *Player) SetLocale(locale string) {
	p.Locale = locale
}

func (p *Player) SetAlreadySendPackets(packets []string) {
	p.AlReadySendPackets = packets
}

func (p *Player) GetAlreadySendPackets() []string {
	return p.AlReadySendPackets
}

func (p *Player) AddAlreadySendPacket(packet string) {
	p.AlReadySendPackets = append(p.AlReadySendPackets, packet)
}

func (p *Player) ContainsAlreadySendPacket(packet string) bool {
	for _, p := range p.AlReadySendPackets {
		if p == packet {
			return true
		}
	}
	return false
}

func (p *Player) GetLocale() string {
	return p.Locale
}

func (p *Player) SetViewDistance(viewDistance byte) {
	p.ViewDistance = viewDistance
}

func (p *Player) GetViewDistance() byte {
	return p.ViewDistance
}

func (p *Player) SetChatMode(chatMode int32) {
	p.ChatMode = chatMode
}

func (p *Player) GetChatMode() int32 {
	return p.ChatMode
}

func (p *Player) SetChatColors(chatColors bool) {
	p.ChatColors = chatColors
}

func (p *Player) GetChatColors() bool {
	return p.ChatColors
}

func (p *Player) SetDisplayedSkinParts(displayedSkinParts byte) {
	p.DisplayedSkinParts = displayedSkinParts
}

func (p *Player) GetDisplayedSkinParts() byte {
	return p.DisplayedSkinParts
}

func (p *Player) SetMainHand(mainHand int32) {
	p.MainHand = mainHand
}

func (p *Player) GetMainHand() int32 {
	return p.MainHand
}

func (p *Player) SetEnableTextFilter(enableTextFilter bool) {
	p.EnableTextFilter = enableTextFilter
}

func (p *Player) GetEnableTextFilter() bool {
	return p.EnableTextFilter
}

func (p *Player) SetAllowServerListings(allowServerListings bool) {
	p.AllowServerListings = allowServerListings
}

func (p *Player) GetAllowServerListings() bool {
	return p.AllowServerListings
}

func (p *Player) SetUUID(newUUID uuid.UUID) {
	p.mu.Lock()
	defer p.mu.Unlock()

	// Construir rutas de los archivos
	basePath, err := os.Getwd()
	if err != nil {
		return
	}
	logDir := filepath.Join(basePath, "logs", "packet")

	oldFilePath := filepath.Join(logDir, fmt.Sprintf("%s.txt", p.UUID))
	newFilePath := filepath.Join(logDir, fmt.Sprintf("%s.txt", newUUID.String()))

	// Abrir el archivo anterior para lectura
	oldFile, err := os.Open(oldFilePath)
	if err != nil {
		if os.IsNotExist(err) {
		} else {
			return
		}
	} else {
		defer oldFile.Close()

		// Abrir el archivo nuevo en modo append
		newFile, err := os.OpenFile(newFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return
		}
		defer newFile.Close()

		// Leer del archivo anterior y escribir en el nuevo
		scanner := bufio.NewScanner(oldFile)
		for scanner.Scan() {
			line := scanner.Text()
			_, err := newFile.WriteString(line + "\n")
			if err != nil {
				return
			}
		}

		if err := scanner.Err(); err != nil {
			return
		}

		// Eliminar el archivo anterior
		err = os.Remove(oldFilePath)
		if err != nil {
			return
		}
	}

	// Actualizar el UUID del jugador
	p.UUID = newUUID
}

func (p *Player) GetUUID() uuid.UUID {
	return p.UUID
}

func NewPlayer(entityId int64, UUID uuid.UUID) *Player {
	return &Player{
		EntityId:        entityId,
		Name:            "Steve",
		Health:          20.0,
		UUID:            UUID,
		PlayerState:     player.Handshake,
		Conn:            nil,
		ProtocolVersion: 0,
		Location:        entities.NewLocation(99, 75, 115, 0, 0),
		OnGround:        false,
		SkinValue:       "ewogICJ0aW1lc3RhbXAiIDogMTcyMDEwMDIxNzMyOCwKICAicHJvZmlsZUlkIiA6ICJiM2E3NjExNGVmMzI0ZjYyYWM4NDRiOWJmNTY1NGFiOSIsCiAgInByb2ZpbGVOYW1lIiA6ICJNcmd1eW1hbnBlcnNvbiIsCiAgInNpZ25hdHVyZVJlcXVpcmVkIiA6IHRydWUsCiAgInRleHR1cmVzIiA6IHsKICAgICJTS0lOIiA6IHsKICAgICAgInVybCIgOiAiaHR0cDovL3RleHR1cmVzLm1pbmVjcmFmdC5uZXQvdGV4dHVyZS81YWFhMzRhMjcxNjY3OGEyMWRmZjU0N2Q0MGUwYjg3MWFmNTNmMzllNzAzYmE5MzMxOGUyNmFhYmFmYWI1MDIwIgogICAgfQogIH0KfQ==",
		SkinSignature:   "h5B2VYWXCVfeIkFx0utRGeDTHGMFyZsGb2I7tbjd6xXp445snJX9XzF4ppxJWnLvTlCvivmOJ+M22hVrV4iqtjXH9AdYXYFspvnflA9fGgNs/dwkDIY6atsgJ8kbmK5EoY1rLU4Dc4w2CrKndVig2cGKvJvWDcOFclu01uNHnbs7F3v/pBeVy6sQA4VtdXUdy5BUGSDD0/M1096TtSJqeeuXzMvHxtDsCGiDGmofhjZDsGAfQvkWbh8zsO4r0tdjoeeP4/32G9AistoZb3Xf98M6m33m2/GuY5T9zGO7WJ3gbA7lvl7qd58bl8yDZAl7LVj3MvMoWG9qvLYnpp1SmKCDYVBK3ZZmq0BFadBm5lbCc5xO1Q7MHJ9hq7Gbf9Z0eYkdQEJ5pL3fXQ2ihsZY+Y5SGqrm0+G40GVz+HQItndpc9mNQcZf1tvPnssSbL+roxrSBG8XfpGz+hIDJkLslkPiLeQRUQPci8sYNHxN8B2otBOc512tOWAdDuYqacj6P1rG6tspH5jYR5q6a5RqTC2i2CoA+sdzq62FV8byQy/pFngJAq+8/svD2WKWAiFUfBcTJP8FZMhQ2AJQW6rq5GqZWb9BPxZV4M3zqDSnjHevgBTwvdWbYBT9yQ8NcrubWHfXEaNUpzdMa41XVBAwSY1smPapK/c2YD60VKffuLg=",
	}
}

func (p *Player) SetProtocolVersion(version int) {
	p.ProtocolVersion = version
}

func (p *Player) GetProtocolVersion() int {
	return p.ProtocolVersion
}

func (p *Player) SetClientState(state player.ClientState) {
	p.PlayerState = state
}

func (p *Player) GetClientState() player.ClientState {
	return p.PlayerState
}

func (p *Player) SetConn(conn net.Conn) {
	p.Conn = &conn
}

func (p *Player) GetConn() *net.Conn {
	return p.Conn
}

func (p *Player) SetName(name string) {
	p.Name = name
}

// GetName devuelve el nombre del jugador
func (p *Player) GetName() string {
	return p.Name
}

// GetHealth devuelve la salud del jugador
func (p *Player) GetHealth() float64 {
	return p.Health
}

func (p *Player) AddPacketLogToPlayer(log string) {
	/*
		TODO: Implementar la escritura de logs en archivos
			go func() {
				basePath, err := os.Getwd()
				if err != nil {
					return
				}
				logDir := filepath.Join(basePath, "logs", "packet")
				logFilePath := filepath.Join(logDir, fmt.Sprintf("%s.txt", p.UUID))

				// Crear los directorios si no existen
				p.mu.Lock() // Evitar condiciones de carrera
				defer p.mu.Unlock()

				err = os.MkdirAll(logDir, os.ModePerm)
				if err != nil {
					return
				}

				// Abrir el archivo en modo append
				file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					return
				}
				defer file.Close()

				// Escribir el log en el archivo
				_, err = file.WriteString(log + "\n")
				if err != nil {
					return
				}
			}()*/
}
