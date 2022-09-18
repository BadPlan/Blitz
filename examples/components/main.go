package main

import (
	"context"
	"github.com/BadPlan/blitz/core"
	"github.com/BadPlan/blitz/core/component/position"
	"github.com/BadPlan/blitz/core/component/size"
	"github.com/BadPlan/blitz/core/component/sprite"
	"github.com/BadPlan/blitz/core/entity"
	"log"
)

func main() {
	err := core.Init()
	if err != nil {
		log.Fatalln(err)
	}

	e1 := entity.New(nil)
	player := entity.New(e1)
	pos1 := &position.Position{X: 200, Y: 300}
	e1.AddComponent(pos1)

	pos := &position.Position{X: 200, Y: 300}
	player.AddComponent(pos)
	spriteComponent, err := sprite.NewSprite("../assets/blitz.png")
	if err != nil {
		log.Fatalln(err)
	}
	player.AddComponent(spriteComponent)
	player.AddComponent(&size.Size{W: 100, H: 200})
	go func(pos *position.Position) {
		for {
			if pos.X >= 1000 {
				pos.X = 0
			}
			pos.X += 0.000001
		}
	}(pos1)

	log.Println(core.MainLoop(context.Background()))
}
