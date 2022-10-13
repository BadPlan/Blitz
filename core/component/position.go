package component

type Position struct {
	X, Y int32
}

func (p *Position) IsComponent() bool {
	return true
}
