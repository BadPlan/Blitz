package component

// LocalToWorld is system component to store data about global position & rotation
type LocalToWorld struct {
	GlobalPosition *Position
	GlobalRotation *Rotation
}

func (l2w *LocalToWorld) IsComponent() bool {
	return true
}
