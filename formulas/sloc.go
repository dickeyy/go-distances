// Package formulas provides implementations of various distance calculation
// formulas for geographical points on a sphere.

package formulas

import (
	"math"

	"github.com/dickeyy/go-distances/utils"
)

// SphericalLawOfCosines calculates the great-circle distance between two points
// using the spherical law of cosines.
//
// Formula is based on:
// https://en.wikipedia.org/wiki/Great-circle_distance
// https://en.wikipedia.org/wiki/Spherical_law_of_cosines
//
// Coordinates are in degrees, and earthRadius is in the desired unit.
func SphericalLawOfCosines(lat1, lon1, lat2, lon2 float64, earthRadius float64) float64 {
	// Convert degrees to radians.
	lat1Rad := utils.DegreeToRad(lat1)
	lon1Rad := utils.DegreeToRad(lon1)
	lat2Rad := utils.DegreeToRad(lat2)
	lon2Rad := utils.DegreeToRad(lon2)

	// Difference in longitude
	deltaLon := lon2Rad - lon1Rad

	// Spherical Law of Cosines formula
	deltaSigma := math.Acos(
		math.Sin(lat1Rad)*math.Sin(lat2Rad) +
			math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Cos(deltaLon),
	)

	// Distance
	distance := float64(earthRadius) * deltaSigma

	return distance
}
