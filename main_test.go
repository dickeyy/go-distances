package main

import (
	"os"
	"slices"
	"testing"
)

// Test data for various scenarios
var testLatitudes = []float64{40.7128, 34.0522, 41.8781}
var testLongitudes = []float64{-74.0060, -118.2437, -87.6298}
var testEarthRadius = 6371.0

func TestCalculateCircularDistanceHaversine(t *testing.T) {
	calculateCircularDistance(testLatitudes, testLongitudes, testEarthRadius, "haversine")
}

func TestCalculateCircularDistanceVincenty(t *testing.T) {
	calculateCircularDistance(testLatitudes, testLongitudes, testEarthRadius, "vincenty")
}

func TestCalculateCircularDistanceSloc(t *testing.T) {
	calculateCircularDistance(testLatitudes, testLongitudes, testEarthRadius, "sloc")
}

func TestCalculateCircularDistanceInvalidFormula(t *testing.T) {
	calculateCircularDistance(testLatitudes, testLongitudes, testEarthRadius, "invalid")
}

func TestCalculateCircularDistanceInsufficientPoints(t *testing.T) {
	calculateCircularDistance([]float64{40.7128}, []float64{-74.0060}, testEarthRadius, "haversine")
}

func TestImportDataFromFileValidJSON(t *testing.T) {
	// Create a temporary test file
	fileContent := `{
		"places": [
			{
				"name": "New York",
				"latitude": "40.7128",
				"longitude": "-74.0060"
			},
			{
				"name": "Los Angeles",
				"latitude": "34.0522",
				"longitude": "-118.2437"
			}
		],
		"earthRadius": 6371.0,
		"formula": "haversine"
	}`

	filePath := "test_valid.json"
	err := os.WriteFile(filePath, []byte(fileContent), 0644)
	if err != nil {
		t.Fatalf("Error creating test file: %v", err)
	}
	defer os.Remove(filePath)

	// Test the function
	importDataFromFile()
}

func TestImportDataFromFileInvalidFormat(t *testing.T) {
	// Create a temporary test file with invalid format
	fileContent := `{
		"places": [
			{
				"name": "New York",
				"latitude": "invalid",
				"longitude": "-74.0060"
			}
		],
		"earthRadius": 6371.0,
		"formula": "haversine"
	}`

	filePath := "test_invalid.json"
	err := os.WriteFile(filePath, []byte(fileContent), 0644)
	if err != nil {
		t.Fatalf("Error creating test file: %v", err)
	}
	defer os.Remove(filePath)

	// Test the function
	importDataFromFile()
}

func TestImportDataFromFileNonJSON(t *testing.T) {
	// Create a temporary test file with non-JSON content
	fileContent := "This is not JSON"
	filePath := "test.txt"
	err := os.WriteFile(filePath, []byte(fileContent), 0644)
	if err != nil {
		t.Fatalf("Error creating test file: %v", err)
	}
	defer os.Remove(filePath)

	// Test the function
	importDataFromFile()
}

func TestImportDataFromFileNonexistent(t *testing.T) {
	// Test with a nonexistent file
	importDataFromFile()
}

func TestValidFormulas(t *testing.T) {
	// Test that all valid formulas are in the slice
	expectedFormulas := []string{"haversine", "vincenty", "sloc"}
	for _, formula := range expectedFormulas {
		if !slices.Contains(validFormulas, formula) {
			t.Errorf("Expected formula %s to be in validFormulas", formula)
		}
	}
}

func TestImportDataFromUser(t *testing.T) {
	importDataFromUser()
}
