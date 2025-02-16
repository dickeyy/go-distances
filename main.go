package main

import (
	"fmt"
	"math"
	"slices"

	"github.com/dickeyy/go-distances/formulas"
)

// calculateCircularDistances computes the distances between points in a circular manner
// using the specified formula. It accepts a variable number of latitudes and longitudes.
func calculateCircularDistance(latitudes []float64, longitudes []float64, earthRadius int, formula string) {
	numPoints := len(latitudes)
	if numPoints < 2 {
		fmt.Println("At least two points are required to calculate circular distances.")
		return
	}

	distances := make([]int, numPoints)

	switch formula {
	case "haversine":
		for i := 0; i < numPoints; i++ {
			nextIndex := (i + 1) % numPoints
			distances[i] = int(math.Round(formulas.Haversine(latitudes[i], longitudes[i], latitudes[nextIndex], longitudes[nextIndex], earthRadius)))
		}
	case "vincenty":
		for i := 0; i < numPoints; i++ {
			nextIndex := (i + 1) % numPoints
			distances[i] = int(math.Round(formulas.Vincenty(latitudes[i], longitudes[i], latitudes[nextIndex], longitudes[nextIndex], earthRadius)))
		}
	default:
		fmt.Println("Invalid formula.")
		return
	}

	fmt.Printf("Circular distances using %s formula:\n", formula)
	for i := 0; i < numPoints; i++ {
		fmt.Printf("Distance %d -> %d: %d units\n", i+1, (i+1)%numPoints+1, distances[i])
	}
}

func main() {
	var numPoints int

	fmt.Print("Enter the number of points: ")
	fmt.Scan(&numPoints)

	latitudes := make([]float64, numPoints)
	longitudes := make([]float64, numPoints)

	fmt.Printf("Enter the latitudes and longitudes of the %d points:\n", numPoints)
	for i := 0; i < numPoints; i++ {
		fmt.Printf("Point %d:\n", i+1)
		fmt.Print("Latitude: ")
		fmt.Scan(&latitudes[i])
		fmt.Print("Longitude: ")
		fmt.Scan(&longitudes[i])
	}

	var earthRadius int
	fmt.Print("Enter the Earth's radius: ")
	fmt.Scan(&earthRadius)

	validFormulas := []string{"haversine", "vincenty"}
	var formula string
	fmt.Print("Enter the formula to use (haversine or vincenty): ")
	fmt.Scan(&formula)

	if !slices.Contains(validFormulas, formula) {
		fmt.Println("Invalid formula.")
		return
	}

	println()
	calculateCircularDistance(latitudes, longitudes, earthRadius, formula)
}
