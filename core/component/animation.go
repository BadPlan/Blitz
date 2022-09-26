package component

import (
	"github.com/BadPlan/blitz/core/utils/image_loader"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

type Animation struct {
	IsActive    bool
	Repeatable  bool
	Duration    time.Duration
	SpreadSheet *sdl.Texture
}

func New(sourcePath string, frameWidth, frameHeight int32, repeatable bool, deltaTime time.Duration) (*Animation, error) {
	var totalW, totalH int32
	im, err := image_loader.LoadImage(sourcePath, &totalW, &totalH)
	if err != nil {
		return nil, err
	}
	return &Animation{
		SpreadSheet: im,
		Repeatable:  repeatable,
		Duration:    deltaTime,
		IsActive:    false,
	}, nil
}

func (a *Animation) IsComponent() bool {
	return true
}
