package sdl2

import (
	"github.com/veandco/go-sdl2/sdl"
)

var Instance SDL2

type SDL2 struct {
	renderer *sdl.Renderer
	window   *sdl.Window
}

func (s *SDL2) SetRenderer(r *sdl.Renderer) {
	s.renderer = r
}

func (s *SDL2) SetWindow(w *sdl.Window) {
	s.window = w
}

func (s *SDL2) GetRenderer() *sdl.Renderer {
	return s.renderer
}

func (s *SDL2) GetWindow() *sdl.Window {
	return s.window
}
