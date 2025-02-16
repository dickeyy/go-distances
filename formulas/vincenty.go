package formulas

import "math"

// Vincenty calculates the great-circle distance between two points on a sphere
// using the given formula. Coordinates are in degrees, and earthRadius is in the desired unit.
func Vincenty(lat1, lon1, lat2, lon2 float64, earthRadius int) float64 {
	// Convert degrees to radians
	rad := math.Pi / 180.0
	lat1Rad := lat1 * rad
	lon1Rad := lon1 * rad
	lat2Rad := lat2 * rad
	lon2Rad := lon2 * rad

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
