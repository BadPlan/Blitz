package main

import (
	"context"
	"github.com/BadPlan/blitz/core"
	"github.com/BadPlan/blitz/core/component"
	"github.com/BadPlan/blitz/core/entity"
	"github.com/BadPlan/blitz/core/utils/cast"
	"log"
	"time"
)

func changeSize(ctx context.Context, sizeParent, sizeChild *component.Size) {
	if sizeParent == nil {
		return
	}
	if sizeChild == nil {
		return
	}
	sign := int32(1)
	for {
		select {
		case <-ctx.Done():
			{
				return
			}
		default:
			{
				if sizeParent.W >= 100 {
					sign = -sign
				} else if sizeParent.W <= 10 {
					sign = -sign
				}
				sizeParent.W, sizeParent.H = sizeParent.W+sign, sizeParent.H+sign
				sizeChild.W, sizeChild.H = sizeChild.W-sign, sizeChild.H-sign
				time.Sleep(time.Millisecond * 5)
			}
		}
	}
}

func main() {
	err := core.Init()
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()
	parent := entity.GetEntityById("0xdddddd")
	parentSize := cast.CastToSize(parent.GetComponent(&component.Size{}))

	child := entity.GetEntityById("0xffffff")
	childSize := cast.CastToSize(child.GetComponent(&component.Size{}))

	go changeSize(ctx, parentSize, childSize)

	log.Println(core.MainLoop(context.Background()))
}
