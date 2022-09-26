package main

import (
	"context"
	"github.com/BadPlan/blitz/core"
	"github.com/BadPlan/blitz/core/component"
	"github.com/BadPlan/blitz/core/entity"
	"log"
	"math"
)

func makeEntities() {
	player := entity.New(nil)
	playerSize := &component.Size{W: 100, H: 100}
	pos := &component.Position{X: 500, Y: 500}
	player.AddComponent(pos)
	rot := &component.Rotation{Angle: math.Pi}

	spriteComponent, err := component.NewSprite("../assets/blitz.png")
	if err != nil {
		log.Fatalln(err)
	}
	player.AddComponent(spriteComponent)
	player.AddComponent(rot)
	player.AddComponent(playerSize)

	ball := entity.New(player)
	ballSize := &component.Size{W: 50, H: 50}
	ballPos := &component.Position{X: 50, Y: 50}
	ballSprite, err := component.NewSprite("../assets/ball.png")
	if err != nil {
		log.Fatalln(err)
	}
	ball.AddComponent(ballPos)
	ball.AddComponent(ballSize)
	ball.AddComponent(ballSprite)
	ballRot := &component.Rotation{Angle: 0}
	ball.AddComponent(ballRot)

	go func(pos *component.Position, rot *component.Rotation, ballRot *component.Rotation) {
		velocity := 0.000001
		for {
			if pos.X >= 1900 {
				pos.X = 1900
				velocity = -velocity
			} else if pos.X <= 0 {
				pos.X = 0
				velocity = -velocity
			}
			pos.X = pos.X + velocity

			rot.Angle += 0.00000001
			if rot.Angle >= 2*math.Pi {
				rot.Angle = 0
			}

			ballRot.Angle += 0.00000001
			if ballRot.Angle >= 2*math.Pi {
				ballRot.Angle = 0
			}
		}
	}(pos, rot, ballRot)
}

func main() {
	err := core.Init()
	if err != nil {
		log.Fatalln(err)
	}
	makeEntities()

	log.Println(core.MainLoop(context.Background()))
}
