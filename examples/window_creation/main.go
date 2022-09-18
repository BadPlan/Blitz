package main

import (
	"context"
	"github.com/BadPlan/blitz/core"
	"log"
)

func main() {
	err := core.Init()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(core.MainLoop(context.Background()))
}
