package entities

type Entity struct {
	ID         int
	Components map[string]interface{}
}

var entityCounter int

func NewEntity() *Entity {
	entityCounter++
	return &Entity{
		ID:         entityCounter,
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

func GetEntityByID(entities []*Entity, id int) *Entity {
	for _, entity := range entities {
		if entity.ID == id {
			return entity
		}
	}
	return nil
}
