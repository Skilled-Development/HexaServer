package manager

import (
	hexaserver_player "HexaServer/entities/player"
	"HexaUtils/entities"
	"HexaUtils/entities/player"
	"net"
	"sync"

	"github.com/google/uuid"
)

var EntityManagerInstance *EntityManager

type EntityManager struct {
	entities sync.Map // Using sync.Map for thread-safe access
	freeIDs  []int64
	nextID   int64
	mu       sync.RWMutex
}

func NewEntityManager() *EntityManager {
	EntityManagerInstance = &EntityManager{
		entities: sync.Map{},
		freeIDs:  make([]int64, 0),
		nextID:   1,
	}
	return EntityManagerInstance
}

// GetFreeID provides a free entity ID
func (em *EntityManager) GetFreeID() int64 {
	em.mu.Lock()
	defer em.mu.Unlock()
	if len(em.freeIDs) > 0 {
		id := em.freeIDs[len(em.freeIDs)-1]
		em.freeIDs = em.freeIDs[:len(em.freeIDs)-1]
		return id
	}
	id := em.nextID
	em.nextID++
	return id
}

// CreatePlayer creates a new player and adds it to the entity manager
func (em *EntityManager) CreatePlayer(conn net.Conn) player.Player {
	id := em.GetFreeID()
	UUID := uuid.New()
	newPlayer := hexaserver_player.NewPlayer(id, UUID)
	newPlayer.SetConn(conn)

	em.entities.Store(UUID, newPlayer) // Store by UUID
	println("Created player with ID", id)

	return newPlayer
}

// RemoveEntity removes an entity from the manager and marks its ID as free
func (em *EntityManager) RemovePlayer(entity player.Player) {
	if entity == nil {
		return
	}

	em.mu.Lock()
	defer em.mu.Unlock()

	em.entities.Delete(entity.GetUUID())

	em.freeIDs = append(em.freeIDs, entity.GetEntityId())

}

// GetEntity retrieves an entity by its UUID
func (em *EntityManager) GetEntity(uuid uuid.UUID) player.Player {
	value, ok := em.entities.Load(uuid)
	if ok {
		return value.(player.Player)
	}
	return nil
}

// GetPlayers returns a slice of all players in the manager
func (em *EntityManager) GetPlayers() []player.Player {
	players := make([]player.Player, 0)
	em.entities.Range(func(key, value interface{}) bool {
		if p, ok := value.(player.Player); ok {
			players = append(players, p)
		}
		return true
	})
	return players
}

// GetPlayersExcept returns a slice of all players except the one with the given entity ID
func (em *EntityManager) GetPlayersExcept(entityID int64) []player.Player {
	players := make([]player.Player, 0)
	em.entities.Range(func(key, value interface{}) bool {
		if p, ok := value.(player.Player); ok {
			if p.GetEntityId() != entityID {
				players = append(players, p)
			}
		}
		return true
	})
	return players
}

// GetAllEntities returns a slice containing all entities managed by the EntityManager.
func (em *EntityManager) GetAllEntities() []entities.Entity {
	allEntities := make([]entities.Entity, 0)
	em.entities.Range(func(key, value interface{}) bool {
		if entity, ok := value.(entities.Entity); ok {
			allEntities = append(allEntities, entity)
		}
		return true
	})
	return allEntities
}

func GetAllEntities() []entities.Entity {
	return EntityManagerInstance.GetAllEntities()
}
