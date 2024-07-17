package entities

import (
	"github.com/google/uuid"
)

type Entity struct {
	ID         uuid.UUID
	Components map[string]interface{}
}

func NewEntity() *Entity {
	return &Entity{
		ID:         uuid.New(),
		Components: make(map[string]interface{}),
	}
}

func (e *Entity) AddComponent(componentID string, component interface{}) {
	e.Components[componentID] = component
}

func (e *Entity) GetComponent(componentID string) interface{} {
	return e.Components[componentID]
}

func (e *Entity) HasComponent(componentID string) bool {
	_, ok := e.Components[componentID]
	return ok
}
