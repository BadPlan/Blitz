package dependency_tree

var Instance Config

const (
	Fullscreen = iota
	Window
)

type (
	Config struct {
		Application struct {
			Name   string `yaml:"name"`
			Screen struct {
				X, Y, W, H int32
				ScreenMode int `yaml:"screen_mode"`
				ClearColor struct {
					R, G, B, A uint8
				} `yaml:"clear_color"`
				FPS int `yaml:"fps"`
			} `yaml:"screen"`
		} `yaml:"application"`
		Entities Entities `yaml:"entities"`
	}
	Entity struct {
		Id         *string    `yaml:"id,omitempty"`
		Components Components `yaml:"components"`
		Children   Entities   `yaml:"children"`
	}
	Entities  []*Entity
	Component struct {
		Type   string `yaml:"type"`
		Fields Fields `yaml:"fields"`
	}
	Components []*Component
	Fields     map[string]interface{}
)
