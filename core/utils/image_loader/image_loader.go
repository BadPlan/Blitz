package image_loader

import (
	"github.com/BadPlan/blitz/core/sdl2"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func LoadImage(path string, W, H *int32) (*sdl.Texture, error) {
	image, err := img.Load(path)
	if err != nil {
		return nil, err
	}
	defer image.Free()
	texture, err := sdl2.Instance.GetRenderer().CreateTextureFromSurface(image)
	if err != nil {
		return nil, err
	}
	*W, *H = image.W, image.H
	return texture, nil
}
