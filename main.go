package main

import (
	"fmt"
	"math"
	"slices"
)

// haversineDistance calculates the great-circle distance between two points
// (lat1, lon1) and (lat2, lon2) using the haversine formula.
// The earthRadius should be provided in the desired unit (e.g., kilometers).
func haversineDistance(lat1, lon1, lat2, lon2 float64, earthRadius int) float64 {
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

// vicinityDistance calculates the great-circle distance using the spherical law of cosines.
// As before, the coordinates are given in degrees.
func vicinityDistance(lat1, lon1, lat2, lon2 float64, earthRadius int) float64 {
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

// calculateCircularDistancesHaversine computes the distances between:
// 1 -> 2, 2 -> 3, 3 -> 4, and 4 -> 1 using the haversine formula.
func calculateCircularDistance(latitudes [4]float64, longitudes [4]float64, earthRadius int, formula string) {
	var d12, d23, d34, d41 int
	switch formula {
	case "haversine":
		d12 = int(math.Round(haversineDistance(latitudes[0], longitudes[0], latitudes[1], longitudes[1], earthRadius)))
		d23 = int(math.Round(haversineDistance(latitudes[1], longitudes[1], latitudes[2], longitudes[2], earthRadius)))
		d34 = int(math.Round(haversineDistance(latitudes[2], longitudes[2], latitudes[3], longitudes[3], earthRadius)))
		d41 = int(math.Round(haversineDistance(latitudes[3], longitudes[3], latitudes[0], longitudes[0], earthRadius)))
	case "vicinity":
		d12 = int(math.Round(vicinityDistance(latitudes[0], longitudes[0], latitudes[1], longitudes[1], earthRadius)))
		d23 = int(math.Round(vicinityDistance(latitudes[1], longitudes[1], latitudes[2], longitudes[2], earthRadius)))
		d34 = int(math.Round(vicinityDistance(latitudes[2], longitudes[2], latitudes[3], longitudes[3], earthRadius)))
		d41 = int(math.Round(vicinityDistance(latitudes[3], longitudes[3], latitudes[0], longitudes[0], earthRadius)))
	}

	fmt.Printf("Circular distances using %s formula:\n", formula)
	fmt.Printf("Distance 1 -> 2: %d units\n", d12)
	fmt.Printf("Distance 2 -> 3: %d units\n", d23)
	fmt.Printf("Distance 3 -> 4: %d units\n", d34)
	fmt.Printf("Distance 4 -> 1: %d units\n", d41)
}

func main() {
	latitudes := [4]float64{75.20479441439075, -63.27864890563778, 61.4327283360288, -2.808585806729539}
	longitudes := [4]float64{-87.42362995032933, -9.284284051915705, 58.393280523530755, 15.367989320561076}

	earthRadius := 6378160000

	// Calculate and print distances using both formulas.
	// get the requested formula from the user input
	validFormulas := []string{"haversine", "vicinity"}
	var formula string
	fmt.Print("Enter the formula to use (haversine or vicinity): ")
	fmt.Scan(&formula)

	if !slices.Contains(validFormulas, formula) {
		fmt.Print("Invalid formula.")
		return
	}

	calculateCircularDistance(latitudes, longitudes, earthRadius, formula)
}
