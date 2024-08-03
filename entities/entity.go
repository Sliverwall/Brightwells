package entities

import (
	"Brightwells/components"
)

type Entity struct {
	ID          int                    // unique id to query entity by
	Components  map[string]interface{} // components map to search for entity's components
	Exist       bool                   // exist boolean to see if entity should be included in entiySlice
	RenderLayer int                    // renderLayer to define draw layers
}

var entityCounter int

func NewEntity(renderLayer int) *Entity {
	entityCounter++
	return &Entity{
		ID:          entityCounter,
		Components:  make(map[string]interface{}),
		Exist:       true,
		RenderLayer: renderLayer,
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

func (e *Entity) KillEntity() {
	e.Exist = false
}

func GetEntityByID(entitySlice []*Entity, id int) *Entity {
	for _, entity := range entitySlice {
		if entity.ID == id {
			return entity
		}
	}
	return nil
}

func GetExistEntitySlice(entitySlice, deadEntitySlice []*Entity) ([]*Entity, []*Entity) {
	var existEntitySlice []*Entity
	var filteredDeadEntitySlice []*Entity

	// Filter entities that exist from entitySlice
	for _, entity := range entitySlice {
		if entity.Exist {
			existEntitySlice = append(existEntitySlice, entity)
		} else {
			deadEntitySlice = append(deadEntitySlice, entity)
		}
	}

	// Filter entities that do not exist from deadEntitySlice
	for _, entity := range deadEntitySlice {
		if !entity.Exist {
			filteredDeadEntitySlice = append(filteredDeadEntitySlice, entity)
		}
	}

	// Check and handle respawn conditions
	existEntitySlice, filteredDeadEntitySlice = EntityRespawn(existEntitySlice, filteredDeadEntitySlice)

	return existEntitySlice, filteredDeadEntitySlice
}

func GetPlayerEntity(entitySlice []*Entity) *Entity {
	for _, entity := range entitySlice {
		if entity.HasComponent(components.PlayerComponentID) {
			return entity
		}
	}
	return nil
}

// NewEntityWithComponents initializes a new entity with given components
func NewEntityWithComponents(layer int, components ...struct {
	ID       string
	Instance interface{}
}) *Entity {
	entity := NewEntity(layer)
	for _, component := range components {
		entity.AddComponent(component.ID, component.Instance)
	}
	return entity
}

// ------------------------------ Respawn -------------------------------

func EntityRespawn(entitySlice, deadEntitySlice []*Entity) ([]*Entity, []*Entity) {

	for _, entity := range deadEntitySlice {
		// DEBUG
		// log.Println("Entity ID in dead slice: ", entity.ID)
		// Check if entity has a spawn point
		if entity.HasComponent(components.SpawnPointComponentID) && !entity.Exist {
			// Check if respawn timer is up
			spawnPoint := entity.GetComponent(components.SpawnPointComponentID).(*components.SpawnPointComponent)
			if spawnPoint.RespawnTime <= spawnPoint.RespawnTimeCount {
				// DEBUG
				// log.Println("Respawned ", entity.ID)
				// Set entity back to alive
				entity.Exist = true
				// Reset respawn timer
				spawnPoint.RespawnTimeCount = 0

				// Set position to spawn positions
				position := entity.GetComponent(components.PositionComponentID).(*components.PositionComponent)

				position.TileX, position.TileY = spawnPoint.TileX, spawnPoint.TileY

				// Set Destination to spawn point as well
				position.DesX, position.DesY = spawnPoint.TileX, spawnPoint.TileY

				// Add entity back to entitySlice
				entitySlice = append(entitySlice, entity)

			} else {
				spawnPoint.RespawnTimeCount += 1 // increase
				// DEBUG
				// log.Println("Spawn counter for ", entity.ID, " ", spawnPoint.RespawnTimeCount)
			}

			// Return entitySlice
			return entitySlice, deadEntitySlice
		} else {
			// Just return entitySlice
			return entitySlice, deadEntitySlice
		}

	}
	return entitySlice, deadEntitySlice
}
