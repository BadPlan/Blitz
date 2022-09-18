package sprite

import (
	"github.com/BadPlan/blitz/core/utils/image_loader"
	"github.com/veandco/go-sdl2/sdl"
)

type Sprite struct {
	Path    string
	W, H    int32
	Texture *sdl.Texture
}

func (s *Sprite) IsComponent() bool {
	return true
}

func NewSprite(path string) (*Sprite, error) {
	var w, h int32
	texture, err := image_loader.LoadImage(path, &w, &h)
	if err != nil {
		return nil, err
	}
	return &Sprite{
		Path:    path,
		Texture: texture,
		W:       w,
		H:       h,
	}, nil
}

func (s *Sprite) GetTexture() *sdl.Texture {
	return s.Texture
}

func (s *Sprite) GetWidth() int32 {
	return s.W
}

func (s *Sprite) GetHeight() int32 {
	return s.H
}
