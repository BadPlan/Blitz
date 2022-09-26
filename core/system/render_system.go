package system

import (
	"context"
	"github.com/BadPlan/blitz/core/component"
	"github.com/BadPlan/blitz/core/dependency_tree"
	"github.com/BadPlan/blitz/core/entity"
	"github.com/BadPlan/blitz/core/errors"
	"github.com/BadPlan/blitz/core/sdl2"
	"github.com/BadPlan/blitz/core/utils/angle"
	"github.com/BadPlan/blitz/core/utils/cast"
	"github.com/BadPlan/blitz/core/utils/fps"
	"github.com/BadPlan/blitz/core/utils/timer"
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

type renderSystem struct {
	timer timer.Timer
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
	rs.timer = timer.NewNonBlocking(fps.DeltaTime(dependency_tree.Instance.Application.Screen.FPS))
	rs.timer.Start()
	return nil
}

func (rs *renderSystem) OnUpdate(ctx context.Context) error {
	for !rs.timer.IsDone() {
	}
	log.Printf("FPS now: %d\n", fps.FramesPerSecond(rs.timer.PrevDelta()))
	rs.timer.Start()
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

func getSpriteParams(e *entity.Entity) (w, h int32, t *sdl.Texture, err error) {
	if e.GetComponent(&component.Sprite{}) == nil {
		err = errors.ErrEmptySpriteComponent
		return
	}
	spriteComponent := cast.CastToSprite(e.GetComponent(&component.Sprite{}))

	w, h = spriteComponent.GetWidth(), spriteComponent.GetHeight()
	t = spriteComponent.GetTexture()
	if t == nil {
		err = errors.ErrEmptySpriteComponent
		return
	}
	return w, h, t, nil
}

func (rs *renderSystem) drawSprite(e *entity.Entity) {
	w, h, t, err := getSpriteParams(e)
	if err != nil {
		log.Println("sprite component is not defined, skip render")
		return
	}
	s := cast.CastToSize(e.GetComponent(&component.Size{}))
	if s == nil {
		log.Println("size component is not defined, skip render")
		return
	}
	l2w := cast.CastToLocalToWorld(e.GetComponent(&component.LocalToWorld{}))
	if l2w == nil {
		log.Println("not specified local to world component")
		return
	}

	pos := l2w.GlobalPosition
	if pos == nil {
		log.Println("global position not specified")
		return
	}

	src := &sdl.Rect{W: w, H: h}
	dst := &sdl.Rect{X: int32(pos.X - s.W/2), Y: int32(pos.Y - s.H/2), W: int32(s.W), H: int32(s.H)}
	renderer := sdl2.Instance.GetRenderer()
	rot := getWorldRotation(e)
	if rot != nil {
		centerX := dst.W / 2
		centerY := dst.H / 2
		err = renderer.CopyEx(t, src, dst, angle.ToDegrees(rot.Angle), &sdl.Point{X: centerX, Y: centerY}, sdl.FLIP_NONE)
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		err = renderer.Copy(t, src, dst)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func newRenderSystem() System {
	return &renderSystem{}
}
