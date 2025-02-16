package main

import (
	"fmt"
	"math"
	"slices"

	"github.com/dickeyy/go-distances/formulas"
)

// calculateCircularDistancesHaversine computes the distances between:
// 1 -> 2, 2 -> 3, 3 -> 4, and 4 -> 1 using the haversine formula.
func calculateCircularDistance(latitudes [4]float64, longitudes [4]float64, earthRadius int, formula string) {
	var d12, d23, d34, d41 int
	switch formula {
	case "haversine":
		d12 = int(math.Round(formulas.Haversine(latitudes[0], longitudes[0], latitudes[1], longitudes[1], earthRadius)))
		d23 = int(math.Round(formulas.Haversine(latitudes[1], longitudes[1], latitudes[2], longitudes[2], earthRadius)))
		d34 = int(math.Round(formulas.Haversine(latitudes[2], longitudes[2], latitudes[3], longitudes[3], earthRadius)))
		d41 = int(math.Round(formulas.Haversine(latitudes[3], longitudes[3], latitudes[0], longitudes[0], earthRadius)))
	case "vicinity":
		d12 = int(math.Round(formulas.Vicinity(latitudes[0], longitudes[0], latitudes[1], longitudes[1], earthRadius)))
		d23 = int(math.Round(formulas.Vicinity(latitudes[1], longitudes[1], latitudes[2], longitudes[2], earthRadius)))
		d34 = int(math.Round(formulas.Vicinity(latitudes[2], longitudes[2], latitudes[3], longitudes[3], earthRadius)))
		d41 = int(math.Round(formulas.Vicinity(latitudes[3], longitudes[3], latitudes[0], longitudes[0], earthRadius)))
	}

	fmt.Printf("Circular distances using %s formula:\n", formula)
	fmt.Printf("Distance 1 -> 2: %d units\n", d12)
	fmt.Printf("Distance 2 -> 3: %d units\n", d23)
	fmt.Printf("Distance 3 -> 4: %d units\n", d34)
	fmt.Printf("Distance 4 -> 1: %d units\n", d41)
}

func main() {
	var latitudes [4]float64
	var longitudes [4]float64
	var earthRadius int
	var formula string

	fmt.Println("Enter the latitudes and longitudes of the four points:")
	for i := 0; i < 4; i++ {
		fmt.Printf("Point %d:\n", i+1)
		fmt.Print("Latitude: ")
		fmt.Scan(&latitudes[i])
		fmt.Print("Longitude: ")
		fmt.Scan(&longitudes[i])
	}

	fmt.Print("Enter the Earth's radius: ")
	fmt.Scan(&earthRadius)

	validFormulas := []string{"haversine", "vicinity"}
	fmt.Print("Enter the formula to use (haversine or vicinity): ")
	fmt.Scan(&formula)

	if !slices.Contains(validFormulas, formula) {
		fmt.Print("Invalid formula.")
		return
	}

	println()
	calculateCircularDistance(latitudes, longitudes, earthRadius, formula)
}
