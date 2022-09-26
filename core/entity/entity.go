package entity

import (
	"github.com/BadPlan/blitz/core/component"
	"github.com/BadPlan/blitz/core/dependency_tree"
	"github.com/BadPlan/blitz/core/utils/id"
	"github.com/modern-go/reflect2"
	"sync"
)

var (
	entities = make(map[int]*Entity)
	mu       sync.Mutex
)

type Entity struct {
	id         int
	parent     *Entity
	components map[component.Component]interface{}
}

func New(parent *Entity) *Entity {
	e := &Entity{
		id:         <-id.Channel,
		parent:     parent,
		components: make(map[component.Component]interface{}),
	}
	e.AddComponent(&component.LocalToWorld{})
	mu.Lock()
	entities[e.id] = e
	mu.Unlock()
	return e
}

func NewBatchedEntities(config dependency_tree.Config) {
	mu.Lock()
	defer mu.Unlock()
	for i := range config.Entities {
		recursiveNewEntity(nil, config.Entities[i])
	}
}

func recursiveNewEntity(parent *Entity, entity *dependency_tree.Entity) {
	// TODO: add entity, go to children
}

func (e *Entity) AddComponent(component component.Component) {
	for i := range e.components {
		if reflect2.TypeOf(component) == reflect2.TypeOf(i) {
			return
		}
	}
	e.components[component] = 0
}
func (e *Entity) GetComponent(component component.Component) *component.Component {
	for i := range e.components {
		if reflect2.TypeOf(component) == reflect2.TypeOf(i) {
			return &i
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

func GetEntityById(id int) *Entity {
	return entities[id]
}

func GetAllEntities() map[int]*Entity {
	return entities
}
