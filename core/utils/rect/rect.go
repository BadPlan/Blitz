package rect

import (
	"github.com/BadPlan/blitz/core/component"
)

func Center(pos component.Position, s component.Size) component.Position {
	center := component.Position{}
	halfX := pos.X + s.W/2
	halfY := pos.Y + s.H/2
	center.X, center.Y = halfX, halfY
	return center
}
