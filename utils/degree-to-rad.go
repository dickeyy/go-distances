// Package utils provides utility functions for the go-distances project,
// including file parsing and degree-to-radian conversion

package utils

import "math"

// DegreeToRad converts an angle from degrees to radians.
func DegreeToRad(degrees float64) float64 {
	return degrees * math.Pi / 180
}
