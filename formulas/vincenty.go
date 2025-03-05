// Package formulas provides implementations of various distance calculation
// formulas for geographical points on a sphere.
package formulas

import (
	"math"

	"github.com/dickeyy/go-distances/utils"
)

// Vincenty calculates the great-circle distance between two points on a sphere
// using a simplified version of the Vincenty formula.
//
// Formula is based on:
// https://en.wikipedia.org/wiki/Great-circle_distance
//
// Coordinates are in degrees, and earthRadius is in the desired unit.
func Vincenty(lat1, lon1, lat2, lon2 float64, earthRadius float64) float64 {
	// Convert degrees to radians
	lat1Rad := utils.DegreeToRad(lat1)
	lon1Rad := utils.DegreeToRad(lon1)
	lat2Rad := utils.DegreeToRad(lat2)
	lon2Rad := utils.DegreeToRad(lon2)

	// Difference in longitude
	deltaLon := lon2Rad - lon1Rad

	// Calculate components of the formula
	A := math.Cos(lat2Rad) * math.Sin(deltaLon)
	B := math.Cos(lat1Rad)*math.Sin(lat2Rad) - math.Sin(lat1Rad)*math.Cos(lat2Rad)*math.Cos(deltaLon)
	C := math.Sin(lat1Rad)*math.Sin(lat2Rad) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Cos(deltaLon)

	// Central angle (Δσ)
	centralAngle := math.Atan2(math.Sqrt(A*A+B*B), C)

	// Distance = radius * central angle
	return float64(earthRadius) * centralAngle
}
