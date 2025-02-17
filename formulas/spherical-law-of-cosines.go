package formulas

import (
	"math"

	"github.com/dickeyy/go-distances/utils"
)

// SphericalLawOfCosines calculates the great-circle distance between two points
// (lat1, lon1) and (lat2, lon2) using the spherical law of cosines.
// The earthRadius should be provided in the desired unit (e.g., kilometers).

// based on the spec from https://en.wikipedia.org/wiki/Spherical_law_of_cosines
// And https://en.wikipedia.org/wiki/Great-circle_distance
func SphericalLawOfCosines(lat1, lon1, lat2, lon2 float64, earthRadius int) float64 {
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
