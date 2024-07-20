package entities

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

func GetExistEntitySlice(entitySlice []*Entity) []*Entity {
	var existEntitySlice []*Entity
	for _, entity := range entitySlice {
		if entity.Exist {
			existEntitySlice = append(existEntitySlice, entity)
		}
	}
	return existEntitySlice
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
