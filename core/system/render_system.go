package system

import (
	"context"
	"github.com/BadPlan/blitz/core/component/size"
	"github.com/BadPlan/blitz/core/component/sprite"
	"github.com/BadPlan/blitz/core/dependency_tree"
	"github.com/BadPlan/blitz/core/entity"
	"github.com/BadPlan/blitz/core/errors"
	"github.com/BadPlan/blitz/core/sdl2"
	"github.com/BadPlan/blitz/core/utils/angle"
	"github.com/BadPlan/blitz/core/utils/fps"
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"time"
)

type renderSystem struct {
	prevFrame time.Time
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
	rs.prevFrame = time.Now()
	return nil
}

func (rs *renderSystem) OnUpdate(ctx context.Context) error {
	passed := time.Now().Sub(rs.prevFrame)
	shouldPass := fps.DeltaTime(dependency_tree.Instance.Application.Screen.FPS)
	if passed < shouldPass {
		<-time.Tick(shouldPass - passed)
	}
	colors := dependency_tree.Instance.Application.Screen.ClearColor
	renderer := sdl2.Instance.GetRenderer()
	renderer.SetDrawColor(colors.R, colors.G, colors.B, colors.A)
	renderer.Clear()
	entities := entity.GetAllEntities()
	for i := range entities {
		rs.drawSprite(entities[i])
	}
	renderer.Present()
	rs.prevFrame = time.Now()
	return nil
}

func getSpriteParams(e *entity.Entity) (w, h int32, t *sdl.Texture, err error) {
	if e.GetComponent(&sprite.Sprite{}) == nil {
		err = errors.ErrEmptySpriteComponent
		return
	}
	spriteComponent := castToSprite(e.GetComponent(&sprite.Sprite{}))

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
	s := castToSize(e.GetComponent(&size.Size{}))
	if s == nil {
		log.Println("size component is not defined, skip render")
		return
	}
	l2w := castToL2W(*GetSystem(&localToWorldSystem{}))
	if l2w == nil {
		log.Println("local_to_world_system is not defined, skip render")
		return
	}

	pos := l2w.PositionByEntityId(e.GetId())
	if e.GetId() == 1 {
		log.Println(pos)
	}

	src := &sdl.Rect{W: w, H: h}
	dst := &sdl.Rect{X: int32(pos.X - s.W/2), Y: int32(pos.Y - s.H/2), W: int32(s.W), H: int32(s.H)}
	renderer := sdl2.Instance.GetRenderer()
	rot := getWorldRotation(e)
	if rot != nil {
		centerX := dst.W / 2
		centerY := dst.H / 2
		if e.GetId() == 1 {
			log.Printf("centerX = %d, centerY = %d, posX = %d, posY = %d", centerX, centerY, int32(pos.X), int32(pos.Y))
		}
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
