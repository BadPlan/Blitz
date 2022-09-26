package cast

import (
	"github.com/BadPlan/blitz/core/component"
)

func CastToPosition(in *component.Component) *component.Position {
	if in == nil {
		return nil
	}
	return asInterface(*in).(*component.Position)
}

func asInterface(component component.Component) interface{} {
	return interface{}(component)
}

func CastToRotation(in *component.Component) *component.Rotation {
	if in == nil {
		return nil
	}
	return asInterface(*in).(*component.Rotation)
}

func CastToSprite(in *component.Component) *component.Sprite {
	if in == nil {
		return nil
	}
	return asInterface(*in).(*component.Sprite)
}

func CastToLocalToWorld(in *component.Component) *component.LocalToWorld {
	if in == nil {
		return nil
	}
	return asInterface(*in).(*component.LocalToWorld)
}

func CastToSize(in *component.Component) *component.Size {
	if in == nil {
		return nil
	}
	return asInterface(*in).(*component.Size)
}
