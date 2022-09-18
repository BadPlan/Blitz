package position

type Position struct {
	X, Y float64
}

func (p *Position) IsComponent() bool {
	return true
}
