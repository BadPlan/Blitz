package component

type Rotation struct {
	Angle float64
}

func (r Rotation) IsComponent() bool {
	return true
}
