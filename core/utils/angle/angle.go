package angle

import "math"

func ToDegrees(a float64) float64 {
	return a * (float64(180) / math.Pi)
}

func ToRadians(a float64) float64 {
	return a * (math.Pi / float64(180))
}
