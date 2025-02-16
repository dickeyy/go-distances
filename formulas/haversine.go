package formulas

import "math"

// haversineDistance calculates the great-circle distance between two points
// (lat1, lon1) and (lat2, lon2) using the haversine formula.
// The earthRadius should be provided in the desired unit (e.g., kilometers).
func Haversine(lat1, lon1, lat2, lon2 float64, earthRadius int) float64 {
	// Convert degrees to radians.
	rad := math.Pi / 180.0
	lat1Rad := lat1 * rad
	lon1Rad := lon1 * rad
	lat2Rad := lat2 * rad
	lon2Rad := lon2 * rad

	// Compute differences.
	dLat := lat2Rad - lat1Rad
	dLon := lon2Rad - lon1Rad

	// Haversine formula.
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return float64(earthRadius) * c
}
