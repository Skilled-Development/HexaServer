package entities

import (
	utils_entities "HexaUtils/entities"

	"github.com/google/uuid"
)

type Entity struct {
	EntityId   int64
	Name       string
	Health     float64
	UUID       uuid.UUID
	Location   *utils_entities.Location
	EntityType utils_entities.EntityType
	WorldID    int64
}

func (e Entity) GetEntityId() int64 {
	return e.EntityId
}

func (e Entity) SetEntityId(id int64) {
	e.EntityId = id
}

func (e Entity) GetName() string {
	return e.Name
}

func (e Entity) SetName(name string) {
	e.Name = name
}

func (e Entity) GetHealth() float64 {
	return e.Health
}

func (e Entity) SetHealth(health float64) {
	e.Health = health
}

func (e Entity) GetUUID() uuid.UUID {
	return e.UUID
}

func (e Entity) SetUUID(uuid uuid.UUID) {
	e.UUID = uuid
}

func (e Entity) GetLocation() *utils_entities.Location {
	return e.Location
}

func (e Entity) SetLocation(location *utils_entities.Location) {
	e.Location = location
}

func (e Entity) GetEntityType() utils_entities.EntityType {
	return e.EntityType
}

func (e Entity) SetEntityType(entityType utils_entities.EntityType) {
	e.EntityType = entityType
}

func (e Entity) GetWorldID() int64 {
	return e.WorldID
}

func (e Entity) SetWorldID(worldID int64) {
	e.WorldID = worldID
}
