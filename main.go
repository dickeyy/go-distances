package main

import (
	"fmt"
	"math"
)

// haversineDistance calculates the great-circle distance between two points
// given their latitude/longitude in degrees using the haversine formula.
// The earthRadius should be provided in the desired unit (e.g., 6371 for kilometers).
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

	// Apply the haversine formula.
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// Return the distance.
	return float64(earthRadius) * c
}

// calculateCircularDistances takes 4 coordinate pairs and calculates:
// 1 -> 2, 2 -> 3, 3 -> 4, and 4 -> 1.
func calculateCircularDistances(lat1, lon1, lat2, lon2, lat3, lon3, lat4, lon4 float64, earthRadius int) {
	d12 := int(math.Round(haversineDistance(lat1, lon1, lat2, lon2, earthRadius)))
	d23 := int(math.Round(haversineDistance(lat2, lon2, lat3, lon3, earthRadius)))
	d34 := int(math.Round(haversineDistance(lat3, lon3, lat4, lon4, earthRadius)))
	d41 := int(math.Round(haversineDistance(lat4, lon4, lat1, lon1, earthRadius)))

	fmt.Printf("Distance between coordinate 1 and coordinate 2: %d units\n", d12)
	fmt.Printf("Distance between coordinate 2 and coordinate 3: %d units\n", d23)
	fmt.Printf("Distance between coordinate 3 and coordinate 4: %d units\n", d34)
	fmt.Printf("Distance between coordinate 4 and coordinate 1: %d units\n", d41)
}

func main() {
	// Example coordinates:
	lat1, lon1 := 75.20479441439075, -87.42362995032933
	lat2, lon2 := -63.27864890563778, -9.284284051915705
	lat3, lon3 := 61.4327283360288, 58.393280523530755
	lat4, lon4 := -2.808585806729539, 15.367989320561076

	// Earth's radius in whatever unit (this is mm)
	earthRadius := 6378160000

	calculateCircularDistances(lat1, lon1, lat2, lon2, lat3, lon3, lat4, lon4, earthRadius)
}
