package system

import (
	"context"
	"github.com/modern-go/reflect2"
)

var systems = []System{newRenderSystem(), newLocalToWorldSystem()}

type System interface {
	OnStart(ctx context.Context) error
	OnUpdate(ctx context.Context) error
}

func GetAllSystems() []System {
	return systems
}

func GetSystem(s System) *System {
	for i := range systems {
		if reflect2.TypeOf(s) == reflect2.TypeOf(systems[i]) {
			return &systems[i]
		}
	}
	return nil
}
