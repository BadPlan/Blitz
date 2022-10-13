package entity

import (
	"github.com/BadPlan/blitz/core/component"
	"github.com/BadPlan/blitz/core/dependency_tree"
	"github.com/BadPlan/blitz/core/utils/id"
	"github.com/modern-go/reflect2"
	"github.com/spf13/cast"
	"sync"
)

var (
	entities = make(map[string]*Entity)
	mu       sync.Mutex
)

type Entity struct {
	id         string
	parent     *Entity
	components map[component.Component]interface{}
}

func newEntity(parent *Entity, identifier *string) *Entity {
	if identifier == nil || *identifier == "" {
		return &Entity{
			id:         <-id.Channel,
			parent:     parent,
			components: make(map[component.Component]interface{}),
		}
	} else {
		return &Entity{
			id:         *identifier,
			parent:     parent,
			components: make(map[component.Component]interface{}),
		}
	}
}

func New(parent *Entity) *Entity {
	e := newEntity(parent, nil)
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

func parseComponent(c *dependency_tree.Component) component.Component {
	switch c.Type {
	case "position":
		{
			return &component.Position{X: cast.ToInt32(c.Fields["x"]), Y: cast.ToInt32(c.Fields["y"])}
		}
	case "size":
		{
			return &component.Size{W: cast.ToInt32(c.Fields["w"]), H: cast.ToInt32(c.Fields["h"])}
		}
	case "sprite":
		{
			s, e := component.NewSprite(c.Fields["path"].(string))
			if e != nil {
				return nil
			}
			return s
		}
	case "rotation":
		{
			return &component.Rotation{Angle: c.Fields["angle"].(float64)}
		}

	default:
		return nil
	}
}

func recursiveNewEntity(parent *Entity, entity *dependency_tree.Entity) {
	node := newEntity(parent, entity.Id)
	for i := range entity.Components {
		parsed := parseComponent(entity.Components[i])
		node.AddComponent(parsed)
	}
	node.AddComponent(&component.LocalToWorld{})
	entities[node.id] = node
	var wg sync.WaitGroup
	wg.Add(len(entity.Children))
	for i := range entity.Children {
		go func(idx int) {
			recursiveNewEntity(node, entity.Children[idx])
			wg.Done()
		}(i)
	}
	wg.Wait()
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

func (e *Entity) GetId() string {
	return e.id
}

func GetEntityById(id string) *Entity {
	return entities[id]
}

func GetAllEntities() map[string]*Entity {
	return entities
}
