// Package utils provides utility functions for the go-distances project,
// including file parsing and degree-to-radian conversion.
package utils

import (
	"encoding/json"
	"os"
	"strconv"
)

type Point struct {
	Name      string `json:"name"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type Data struct {
	Places      []Point `json:"places"`
	EarthRadius float64 `json:"earthRadius"`
	Formula     string  `json:"formula"`
}

// ParseFile reads a JSON file containing geographical point data and returns
// the parsed information.
//
// It returns the number of points, slices of latitudes and longitudes,
// the Earth's radius, and the formula to use for distance calculation.
func ParseFile(filePath string) (numPoints int, latitudes []float64, longitudes []float64, earthRadius float64, formula string, err error) {
	// read the file
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	// decode the file
	decoder := json.NewDecoder(file)
	data := Data{}
	err = decoder.Decode(&data)
	if err != nil {
		return
	}

	numPoints = len(data.Places)
	latitudes = make([]float64, numPoints)
	longitudes = make([]float64, numPoints)
	for i := range numPoints {
		latitudes[i], err = strconv.ParseFloat(data.Places[i].Latitude, 64)
		if err != nil {
			return
		}
		longitudes[i], err = strconv.ParseFloat(data.Places[i].Longitude, 64)
		if err != nil {
			return
		}
	}
	earthRadius = data.EarthRadius
	formula = data.Formula

	// if formula does not exist, default to Vincenty
	if formula == "" {
		formula = "vincenty"
	} else if formula == "spherical law of cosines" {
		formula = "sloc"
	}

	return
}
