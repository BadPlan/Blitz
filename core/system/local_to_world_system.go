package system

import (
	"context"
	"github.com/BadPlan/blitz/core/component"
	"github.com/BadPlan/blitz/core/entity"
	"github.com/BadPlan/blitz/core/utils/cast"
	"log"
	"math"
)

type localToWorldSystem struct {
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
	return &localToWorldSystem{}
}

func (l2w *localToWorldSystem) processEntity(e *entity.Entity) {
	pos := getWorldPosition(e)
	c := cast.CastToLocalToWorld(e.GetComponent(&component.LocalToWorld{}))
	if c == nil {
		c = new(component.LocalToWorld)
		e.AddComponent(c)
	}
	if pos != nil {
		c.GlobalPosition = pos
	}
	rot := getWorldRotation(e)
	if rot != nil {
		c.GlobalRotation = rot
	}
}

func getWorldRotation(e *entity.Entity) *component.Rotation {
	parent := e.GetParent()
	if parent == nil {
		return cast.CastToRotation(e.GetComponent(&component.Rotation{}))
	}
	rot := cast.CastToRotation(parent.GetComponent(&component.Rotation{}))
	if rot == nil {
		return &component.Rotation{Angle: 0}
	}
	return &component.Rotation{Angle: rot.Angle + cast.CastToRotation(e.GetComponent(&component.Rotation{})).Angle}
}

func getWorldPosition(e *entity.Entity) *component.Position {
	parent := e.GetParent()
	if parent == nil {
		// log.Printf("id %s, nil parent\n", e.GetId())
		pos := cast.CastToPosition(e.GetComponent(&component.Position{}))
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
	parentSize := cast.CastToSize(parent.GetComponent(&component.Size{}))
	if parentSize == nil {
		parentSize = &component.Size{}
	}
	localSize := cast.CastToSize(e.GetComponent(&component.Size{}))
	if localSize == nil {
		localSize = &component.Size{}
	}
	localPos := cast.CastToPosition(e.GetComponent(&component.Position{}))
	if localPos == nil {
		log.Println("nil local pos")
		return nil
	}
	parentRot := cast.CastToRotation(parent.GetComponent(&component.Rotation{}))
	parentCenter := component.Position{X: posParent.X + parentSize.W/2, Y: posParent.Y + parentSize.H/2}
	if parentRot != nil {
		length := math.Sqrt(float64(localPos.X*localPos.X + localPos.Y*localPos.Y))
		l := component.Position{
			X: int32(length*math.Cos(parentRot.Angle) - length*math.Sin(parentRot.Angle)),
			Y: int32(length*math.Sin(parentRot.Angle) + length*math.Cos(parentRot.Angle)),
		}
		return &component.Position{X: parentCenter.X + l.X - localSize.W/2, Y: parentCenter.Y + l.Y - localSize.H/2}
	}
	return &component.Position{X: parentCenter.X + localPos.X, Y: parentCenter.Y + localPos.Y}
}
