package entity

import (
	"github.com/BadPlan/blitz/core/component"
	"github.com/BadPlan/blitz/core/utils/id"
	"github.com/modern-go/reflect2"
)

var entities []*Entity

type Entity struct {
	id         int
	parent     *Entity
	components []component.Component
}

func New(parent *Entity) *Entity {
	e := &Entity{
		id:         <-id.Channel,
		parent:     parent,
		components: nil,
	}
	entities = append(entities, e)
	return e
}

func (e *Entity) AddComponent(component component.Component) {
	for i := range e.components {
		if reflect2.TypeOf(component) == reflect2.TypeOf(e.components[i]) {
			return
		}
	}
	e.components = append(e.components, component)
}
func (e *Entity) GetComponent(component component.Component) *component.Component {
	for i := range e.components {
		if reflect2.TypeOf(component) == reflect2.TypeOf(e.components[i]) {
			return &e.components[i]
		}
	}
	return nil
}

func (e *Entity) GetParent() *Entity {
	return e.parent
}

func (e *Entity) GetId() int {
	return e.id
}

func GetAllEntities() []*Entity {
	return entities
}
