package system

import (
	"context"
)

type DependencySystem struct {
}

func (ds *DependencySystem) OnStart(ctx context.Context) error {
	//list := dependency_tree.Instance.Entities
	//for i := range list {
	//	entity.NewParsedEntity(list[i].Parent)
	//}
	return nil
}

func (ds *DependencySystem) OnUpdate(ctx context.Context) error {
	return nil
}
