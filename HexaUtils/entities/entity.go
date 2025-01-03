package entities

import "github.com/google/uuid"

type Entity interface {
	GetEntityId() int64
	GetName() string
	GetHealth() float64
	GetUUID() uuid.UUID
	GetLocation() *Location
	GetEntityType() EntityType
	GetWorldID() int64
}
