package rect

import (
	"github.com/BadPlan/blitz/core/component/position"
	"github.com/BadPlan/blitz/core/component/size"
)

func Center(pos position.Position, s size.Size) position.Position {
	center := position.Position{}
	halfX := pos.X + s.W/2
	halfY := pos.Y + s.H/2
	center.X, center.Y = halfX, halfY
	return center
}
