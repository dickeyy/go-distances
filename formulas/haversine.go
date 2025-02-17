package formulas

import (
	"math"

	"github.com/dickeyy/go-distances/utils"
)

// Haversine calculates the great-circle distance between two points
// (lat1, lon1) and (lat2, lon2) using the haversine formula.
// The earthRadius should be provided in the desired unit (e.g., kilometers).

// based on the spec from https://en.wikipedia.org/wiki/Haversine_formula
// And https://en.wikipedia.org/wiki/Great-circle_distance
func Haversine(lat1, lon1, lat2, lon2 float64, earthRadius int) float64 {
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
