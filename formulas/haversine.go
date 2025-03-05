// Package formulas provides implementations of various distance calculation
// formulas for geographical points on a sphere.

package formulas

import (
	"math"

	"github.com/dickeyy/go-distances/utils"
)

// Haversine calculates the great-circle distance between two points
// using the haversine formula.
//
// Formula is based on:
// https://en.wikipedia.org/wiki/Great-circle_distance
// https://en.wikipedia.org/wiki/Haversine_formula
//
// Coordinates are in degrees, and earthRadius is in the desired unit.
func Haversine(lat1, lon1, lat2, lon2 float64, earthRadius float64) float64 {
	// Convert degrees to radians.
	lat1Rad := utils.DegreeToRad(lat1)
	lon1Rad := utils.DegreeToRad(lon1)
	lat2Rad := utils.DegreeToRad(lat2)
	lon2Rad := utils.DegreeToRad(lon2)

	// Differences in coordinates
	deltaLat := lat2Rad - lat1Rad
	deltaLon := lon2Rad - lon1Rad

	// Haversine formula
	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(deltaLon/2)*math.Sin(deltaLon/2)
	c := 2 * math.Asin(math.Sqrt(a))

	// Distance
	distance := float64(earthRadius) * c

	return distance
}
