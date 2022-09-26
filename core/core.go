package core

import (
	"context"
	"github.com/BadPlan/blitz/core/dependency_tree"
	"github.com/BadPlan/blitz/core/sdl2"
	"github.com/BadPlan/blitz/core/system"
	"github.com/spf13/viper"
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"os"
)

func Init() error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	v := viper.New()
	v.AddConfigPath(path)
	err = v.ReadInConfig()
	err = v.Unmarshal(&dependency_tree.Instance)
	if err != nil {
		log.Println(err)
		return err
	}
	err = sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		log.Println(err)
		return err
	}
	for i := range system.GetAllSystems() {
		err = system.GetAllSystems()[i].OnStart(context.Background())
		if err != nil {
			return err
		}
	}
	return nil
}

func MainLoop(ctx context.Context) error {
	var (
		err error
	)
	for {
		for ev := sdl.PollEvent(); ev != nil; ev = sdl.PollEvent() {
			switch ev.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				err = Terminate()
				return err
			}
		}
		for i := range system.GetAllSystems() {
			err = system.GetAllSystems()[i].OnUpdate(ctx)
			if err != nil {
				return err
			}
		}
	}
}

func Terminate() error {
	err := sdl2.Instance.GetWindow().Destroy()
	sdl.Quit()
	return err
}
