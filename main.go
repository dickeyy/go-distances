package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/dickeyy/go-distances/formulas"
	"github.com/dickeyy/go-distances/utils"
)

var numPoints int
var importFile string
var formula string
var earthRadius int
var latitudes []float64
var longitudes []float64

var validFormulas = []string{"haversine", "vincenty"}

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

	fmt.Printf("\nCircular distances using %s formula:\n", formula)
	for i := 0; i < numPoints; i++ {
		fmt.Printf("Distance %d -> %d: %d units\n", i+1, (i+1)%numPoints+1, distances[i])
	}
}

func importDataFromUser() {
	fmt.Print("Enter the number of points: ")
	fmt.Scan(&numPoints)

	fmt.Printf("Enter the latitudes and longitudes of the %d points:\n", numPoints)
	for i := 0; i < numPoints; i++ {
		fmt.Printf("Point %d:\n", i+1)
		fmt.Print("Latitude: ")
		fmt.Scan(&latitudes[i])
		fmt.Print("Longitude: ")
		fmt.Scan(&longitudes[i])
	}

	fmt.Print("Enter the Earth's radius: ")
	fmt.Scan(&earthRadius)

	fmt.Print("Enter the formula to use (haversine or vincenty): ")
	fmt.Scan(&formula)

	if !slices.Contains(validFormulas, formula) {
		fmt.Println("Invalid formula.")
		return
	}
}

func importDataFromFile() {
	fmt.Print("Enter the path to the file: ")
	var filePath string
	fmt.Scan(&filePath)

	// make sure the file is json
	if !strings.HasSuffix(filePath, ".json") {
		fmt.Println("Invalid file format. Please enter a JSON file.")
		return
	}

	// read the file
	var err error
	numPoints, latitudes, longitudes, earthRadius, formula, err = utils.ParseFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	if !slices.Contains(validFormulas, formula) {
		fmt.Println("Invalid formula.")
		return
	}

	fmt.Printf("Data imported from %s:\n", filePath)
}

func main() {
	fmt.Print("Do you want to import points from a file? (y/n): ")
	fmt.Scan(&importFile)

	if importFile == "y" {
		importDataFromFile()
	} else {
		importDataFromUser()
	}

	calculateCircularDistance(latitudes, longitudes, earthRadius, formula)
}
