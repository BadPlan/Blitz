package system

import (
	"context"
	"github.com/BadPlan/blitz/core/component/size"
	"github.com/BadPlan/blitz/core/component/sprite"
	"github.com/BadPlan/blitz/core/dependency_tree"
	"github.com/BadPlan/blitz/core/entity"
	"github.com/BadPlan/blitz/core/sdl2"
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

type renderSystem struct {
}

func (rs *renderSystem) OnStart(ctx context.Context) error {
	screen := dependency_tree.Instance.Application.Screen
	var screenMode uint32
	if screen.ScreenMode == 0 {
		screenMode = sdl.WINDOW_FULLSCREEN | sdl.WINDOW_SHOWN
	} else {
		screenMode = sdl.WINDOW_SHOWN
	}
	win, err := sdl.CreateWindow(
		dependency_tree.Instance.Application.Name,
		screen.X,
		screen.Y,
		screen.W,
		screen.H,
		screenMode,
	)
	if err != nil {
		return err
	}
	sdl2.Instance.SetWindow(win)
	renderer, err := sdl.CreateRenderer(win, 0, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return err
	}
	sdl2.Instance.SetRenderer(renderer)
	return nil
}

func (rs *renderSystem) OnUpdate(ctx context.Context) error {
	colors := dependency_tree.Instance.Application.Screen.ClearColor
	renderer := sdl2.Instance.GetRenderer()
	renderer.SetDrawColor(colors.R, colors.G, colors.B, colors.A)
	renderer.Clear()
	entities := entity.GetAllEntities()
	for i := range entities {
		rs.drawSprite(entities[i])
	}
	renderer.Present()
	return nil
}

func (rs *renderSystem) drawSprite(e *entity.Entity) {
	if e.GetComponent(&sprite.Sprite{}) == nil {
		return
	}
	spriteComponent := castToSprite(*e.GetComponent(&sprite.Sprite{}))

	w, h := spriteComponent.GetWidth(), spriteComponent.GetHeight()
	t := spriteComponent.GetTexture()
	if t == nil {
		return
	}
	s := castToSize(*e.GetComponent(&size.Size{}))
	if s == nil {
		return
	}
	l2w := castToL2W(*GetSystem(&localToWorldSystem{}))
	if l2w == nil {
		log.Println("local_to_world_system is not defined, skip render")
		return
	}

	pos := l2w.PositionByEntityId(e.GetId())

	src := &sdl.Rect{W: w, H: h}
	dst := &sdl.Rect{X: int32(pos.X), Y: int32(pos.Y), W: int32(s.W), H: int32(s.H)}

	renderer := sdl2.Instance.GetRenderer()
	err := renderer.Copy(t, src, dst)
	if err != nil {
		log.Println(err)
		return
	}
}

func castToSprite(in interface{}) *sprite.Sprite {
	return in.(*sprite.Sprite)
}

func castToL2W(in interface{}) *localToWorldSystem {
	return in.(*localToWorldSystem)
}

func castToSize(in interface{}) *size.Size {
	return in.(*size.Size)
}

func newRenderSystem() System {
	return &renderSystem{}
}
