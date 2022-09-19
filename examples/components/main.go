package main

import (
	"context"
	"github.com/BadPlan/blitz/core"
	"github.com/BadPlan/blitz/core/component/position"
	"github.com/BadPlan/blitz/core/component/rotation"
	"github.com/BadPlan/blitz/core/component/size"
	"github.com/BadPlan/blitz/core/component/sprite"
	"github.com/BadPlan/blitz/core/entity"
	"log"
	"math"
)

func makeEntities() {
	player := entity.New(nil)
	playerSize := &size.Size{W: 100, H: 100}
	pos := &position.Position{X: 500, Y: 500}
	player.AddComponent(pos)
	rot := &rotation.Rotation{Angle: math.Pi}

	spriteComponent, err := sprite.NewSprite("../assets/blitz.png")
	if err != nil {
		log.Fatalln(err)
	}
	player.AddComponent(spriteComponent)
	player.AddComponent(rot)
	player.AddComponent(playerSize)

	ball := entity.New(player)
	ballSize := &size.Size{W: 20, H: 20}
	ballPos := &position.Position{X: 50, Y: 10}
	ballSprite, err := sprite.NewSprite("../assets/ball.png")
	if err != nil {
		log.Fatalln(err)
	}
	ball.AddComponent(ballPos)
	ball.AddComponent(ballSize)
	ball.AddComponent(ballSprite)

	go func(pos *position.Position, rot *rotation.Rotation) {
		for {
			//if pos.X >= 1000 {
			//	pos.X = 0
			//}
			//pos.X += 0.000001
			rot.Angle += 0.00000001
			if rot.Angle >= 2*math.Pi {
				rot.Angle = 0
			}
		}
	}(pos, rot)
}

func main() {
	err := core.Init()
	if err != nil {
		log.Fatalln(err)
	}
	makeEntities()

	log.Println(core.MainLoop(context.Background()))
}
