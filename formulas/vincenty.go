package formulas

import "math"

// vincenty calculates the great-circle distance using the spherical law of cosines.
// As before, the coordinates are given in degrees.
func Vincenty(lat1, lon1, lat2, lon2 float64, earthRadius int) float64 {
	// Convert degrees to radians.
	rad := math.Pi / 180.0
	lat1Rad := lat1 * rad
	lon1Rad := lon1 * rad
	lat2Rad := lat2 * rad
	lon2Rad := lon2 * rad

	// Calculate the cosine of the central angle.
	cosDelta := math.Sin(lat1Rad)*math.Sin(lat2Rad) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Cos(lon2Rad-lon1Rad)

	// Clamp the cosine between -1 and 1 to avoid numerical issues.
	if cosDelta > 1 {
		cosDelta = 1
	} else if cosDelta < -1 {
		cosDelta = -1
	}

	centralAngle := math.Acos(cosDelta)
	return float64(earthRadius) * centralAngle
}
