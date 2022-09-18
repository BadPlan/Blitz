package dependency_tree

var Instance Config

const (
	Fullscreen = iota
	Window
)

type Config struct {
	Application struct {
		Name   string
		Screen struct {
			X, Y, W, H int32
			ScreenMode int
			ClearColor struct {
				R, G, B, A uint8
			}
		}
	}
}
