package size

type Size struct {
	H, W float64
}

func (s *Size) IsComponent() bool {
	return true
}
