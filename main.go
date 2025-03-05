// Package main provides a program for calculating great-circle distances between
// geographical points using various formulas.
//
// The program supports three distance calculation methods:
//   - Haversine formula
//   - Vincenty formula (simplified version)
//   - Spherical Law of Cosines (SLOC)
//
// Users can input data manually or from a JSON file, specify the Earth's radius,
// and choose the calculation formula.
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
var earthRadius float64
var latitudes []float64
var longitudes []float64

var validFormulas = []string{"haversine", "vincenty", "sloc"}

// calculateCircularDistance computes the distances between points in a circular manner
// using the specified formula. It accepts slices of latitudes and longitudes,
// the Earth's radius, and the formula name.
//
// Supported formulas are "haversine", "vincenty", and "sloc".
//
// The function prints the calculated distances between consecutive points,
// wrapping around to the first point after the last one.
func calculateCircularDistance(latitudes []float64, longitudes []float64, earthRadius float64, formula string) {
	numPoints := len(latitudes)
	if numPoints < 2 {
		fmt.Println("At least two points are required to calculate circular distances.")
		return
	}

	distances := make([]int, numPoints)

	switch formula {
	case "haversine":
		for i := range numPoints {
			nextIndex := (i + 1) % numPoints
			distances[i] = int(math.Round(formulas.Haversine(latitudes[i], longitudes[i], latitudes[nextIndex], longitudes[nextIndex], earthRadius)))
		}
	case "vincenty":
		for i := range numPoints {
			nextIndex := (i + 1) % numPoints
			distances[i] = int(math.Round(formulas.Vincenty(latitudes[i], longitudes[i], latitudes[nextIndex], longitudes[nextIndex], earthRadius)))
		}
	case "sloc":
		for i := range numPoints {
			nextIndex := (i + 1) % numPoints
			distances[i] = int(math.Round(formulas.SphericalLawOfCosines(latitudes[i], longitudes[i], latitudes[nextIndex], longitudes[nextIndex], earthRadius)))
		}
	default:
		fmt.Println("Invalid formula.")
		return
	}

	fmt.Printf("\nCircular distances using %s formula:\n", formula)
	for i := range numPoints {
		fmt.Printf("Distance %d -> %d: %d units\n", i+1, (i+1)%numPoints+1, distances[i])
	}
}

// importDataFromUser prompts the user to enter the number of points,
// their latitudes and longitudes, the Earth's radius, and the formula to use.
// It populates the global variables with the input data.
func importDataFromUser() {
	fmt.Print("Enter the number of points: ")
	fmt.Scan(&numPoints)

	latitudes = make([]float64, numPoints)
	longitudes = make([]float64, numPoints)

	fmt.Printf("Enter the latitudes and longitudes of the %d points:\n", numPoints)
	for i := range numPoints {
		fmt.Printf("Point %d:\n", i+1)
		fmt.Print("Latitude: ")
		fmt.Scan(&latitudes[i])
		fmt.Print("Longitude: ")
		fmt.Scan(&longitudes[i])
	}

	fmt.Print("Enter the Earth's radius: ")
	fmt.Scan(&earthRadius)

	fmt.Print("Enter the formula to use (haversine or vincenty or sloc): ")
	fmt.Scan(&formula)

	if !slices.Contains(validFormulas, formula) {
		fmt.Println("Invalid formula.")
		return
	}
}

// importDataFromFile prompts the user for a JSON file path and reads the data
// from the file. It populates the global variables with the imported data.
//
// The JSON file should contain an array of points with latitudes and longitudes,
// the Earth's radius, and the formula to use.
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

// main is the entry point of the program. It asks the user whether to import
// data from a file or enter it manually, then calculates and displays the
// circular distances between the points using the specified formula.
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
