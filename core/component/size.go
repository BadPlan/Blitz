package component

type Size struct {
	H, W int32
}

func (s *Size) IsComponent() bool {
	return true
}
