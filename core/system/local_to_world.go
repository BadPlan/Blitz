package system

import (
	"context"
	"github.com/BadPlan/blitz/core/component"
	"github.com/BadPlan/blitz/core/component/position"
	"github.com/BadPlan/blitz/core/component/rotation"
	"github.com/BadPlan/blitz/core/component/size"
	"github.com/BadPlan/blitz/core/component/sprite"
	"github.com/BadPlan/blitz/core/entity"
	"log"
	"math"
)

type localToWorldSystem struct {
	positions map[int]position.Position
	rotations map[int]rotation.Rotation
}

func (l2w *localToWorldSystem) OnStart(ctx context.Context) error {
	return nil
}

func (l2w *localToWorldSystem) OnUpdate(ctx context.Context) error {
	entities := entity.GetAllEntities()
	for i := range entities {
		l2w.processEntity(entities[i])
	}
	return nil
}

func newLocalToWorldSystem() System {
	return &localToWorldSystem{
		positions: make(map[int]position.Position),
		rotations: make(map[int]rotation.Rotation),
	}
}

func (l2w *localToWorldSystem) processEntity(e *entity.Entity) {
	pos := getWorldPosition(e)
	if pos != nil {
		l2w.positions[e.GetId()] = *pos
	}
	rot := getWorldRotation(e)
	if rot != nil {
		l2w.rotations[e.GetId()] = *rot
	}
}

func getWorldRotation(e *entity.Entity) *rotation.Rotation {
	parent := e.GetParent()
	if parent == nil {
		return castToRotation(e.GetComponent(&rotation.Rotation{}))
	}
	rot := castToRotation(parent.GetComponent(&rotation.Rotation{}))
	return rot
}

func getWorldPosition(e *entity.Entity) *position.Position {
	parent := e.GetParent()
	if parent == nil {
		log.Printf("id %d, nil parent\n", e.GetId())
		pos := castToPosition(e.GetComponent(&position.Position{}))
		if pos == nil {
			return nil
		}
		return pos
	}
	posParent := getWorldPosition(e.GetParent())
	if posParent == nil {
		log.Printf("id %d, nil parent pos\n", e.GetId())
		return nil
	}
	parentSize := castToSize(parent.GetComponent(&size.Size{}))
	if parentSize == nil {
		parentSize = &size.Size{}
	}
	localSize := castToSize(parent.GetComponent(&size.Size{}))
	if localSize == nil {
		localSize = &size.Size{}
	}
	localPos := castToPosition(e.GetComponent(&position.Position{}))
	if localPos == nil {
		log.Println("nil local pos")
		return nil
	}
	parentRot := castToRotation(parent.GetComponent(&rotation.Rotation{}))
	parentCenter := position.Position{X: posParent.X + parentSize.W/2, Y: posParent.Y + parentSize.H/2}
	//localCenter := position.Position{X: localPos.X - localSize.W, Y: localPos.Y - localSize.H}
	if parentRot != nil {
		length := math.Sqrt(localPos.X*localPos.X + localPos.Y*localPos.Y)
		log.Printf("length %f", length)
		l := position.Position{
			X: length*math.Cos(parentRot.Angle) - length*math.Sin(parentRot.Angle),
			Y: length*math.Sin(parentRot.Angle) + length*math.Cos(parentRot.Angle),
		}
		log.Printf("parent center = %f, %f", parentCenter.X, parentCenter.Y)
		log.Printf("l := %f,%f", l.X, l.Y)
		return &position.Position{X: parentCenter.X + l.X - localSize.W/2, Y: parentCenter.Y + l.Y - localSize.H/2}
	}
	return &position.Position{X: localPos.X, Y: localPos.Y}
}

func castToPosition(in *component.Component) *position.Position {
	if in == nil {
		return nil
	}
	return asInterface(*in).(*position.Position)
}

func asInterface(component component.Component) interface{} {
	return interface{}(component)
}

func castToRotation(in *component.Component) *rotation.Rotation {
	if in == nil {
		return nil
	}
	return asInterface(*in).(*rotation.Rotation)
}

func castToSprite(in *component.Component) *sprite.Sprite {
	if in == nil {
		return nil
	}
	return asInterface(*in).(*sprite.Sprite)
}

func castToL2W(in interface{}) *localToWorldSystem {
	return in.(*localToWorldSystem)
}

func castToSize(in *component.Component) *size.Size {
	if in == nil {
		return nil
	}
	return asInterface(*in).(*size.Size)
}

func (l2w *localToWorldSystem) PositionByEntityId(id int) position.Position {
	return l2w.positions[id]
}

func Length(pos1, pos2 position.Position) float64 {
	deltaX := pos1.X - pos2.X
	deltaY := pos1.Y - pos2.Y
	return math.Sqrt(math.Pow(deltaX, 2) + math.Pow(deltaY, 2))
}
