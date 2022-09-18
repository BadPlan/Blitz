package system

import (
	"context"
	"github.com/BadPlan/blitz/core/component/position"
	"github.com/BadPlan/blitz/core/entity"
)

type localToWorldSystem struct {
	positions map[int]position.Position
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
	return &localToWorldSystem{positions: make(map[int]position.Position)}
}

func (l2w *localToWorldSystem) processEntity(e *entity.Entity) {
	pos := getWorldPosition(e)
	l2w.positions[e.GetId()] = pos
}

func getWorldPosition(e *entity.Entity) position.Position {
	if e.GetParent() == nil {
		pos := castToPosition(*e.GetComponent(&position.Position{}))
		if pos == nil {
			return position.Position{X: 0, Y: 0}
		}
		return *castToPosition(pos)
	}
	posPrev := getWorldPosition(e.GetParent())
	local := castToPosition(*e.GetComponent(&position.Position{}))
	return position.Position{X: posPrev.X + local.X, Y: posPrev.Y + local.Y}
}

func castToPosition(in interface{}) *position.Position {
	if in == nil {
		return nil
	}
	return in.(*position.Position)
}

func (l2w *localToWorldSystem) PositionByEntityId(id int) position.Position {
	return l2w.positions[id]
}
