package utils

import (
	"os"
	"testing"
)

var validFile2PointsVincenty = `{
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
		"formula": "vincenty"
	}`

var validFileSloc = `{
		"places": [],
		"earthRadius": 6371.0,
		"formula": "sloc"
	}`

var validFileSphericalLawOfCosines = `{
		"places": [],
		"earthRadius": 6371.0,
		"formula": "spherical law of cosines"
	}`

var validFileHaversine = `{
		"places": [],	
		"earthRadius": 6371.0,
		"formula": "haversine"
	}`

var emptyFormulaFile = `{
		"places": [],
		"earthRadius": 6371.0
	}`

var invalidJSONFile = `{
		places: [],
	}`

var invalidLatitudeFile = `{
		"places": [
			{
				"name": "New York",
				"latitude": "invalid",
				"longitude": "-74.0060"
			}
		],
		"earthRadius": 6371.0,
		"formula": "vincenty"
	}`

var invalidLongitudeFile = `{
		"places": [
			{
				"name": "New York",
				"latitude": "40.7128",
				"longitude": "invalid"
			}
		],
		"earthRadius": 6371.0,	
		"formula": "vincenty"
	}`

func makeTestFile(data string) (string, error) {
	filePath := "test.json"
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// write the test data to the file
	file.WriteString(data)

	return filePath, nil
}

func TestParseFile(t *testing.T) {
	filePath, err := makeTestFile(validFile2PointsVincenty)
	if err != nil {
		t.Fatalf("Error creating test file: %v", err)
	}
	defer os.Remove(filePath)

	// read the test data from the file
	numPoints, latitudes, longitudes, earthRadius, formula, err := ParseFile(filePath)
	if err != nil {
		t.Fatalf("Error parsing file: %v", err)
	}

	// check the results
	if numPoints != 2 {
		t.Fatalf("Expected 2 points, got %d", numPoints)
	}
	if latitudes[0] != 40.7128 {
		t.Fatalf("Expected latitude 40.7128, got %f", latitudes[0])
	}
	if longitudes[0] != -74.0060 {
		t.Fatalf("Expected longitude -74.0060, got %f", longitudes[0])
	}
	if earthRadius != 6371.0 {
		t.Fatalf("Expected earth radius 6371.0, got %f", earthRadius)
	}
	if formula != "vincenty" {
		t.Fatalf("Expected formula vincenty, got %s", formula)
	}
}

func TestParseFileSloc(t *testing.T) {
	filePath, err := makeTestFile(validFileSloc)
	if err != nil {
		t.Fatalf("Error creating test file: %v", err)
	}
	defer os.Remove(filePath)

	// read the test data from the file
	_, _, _, _, formula, err := ParseFile(filePath)
	if err != nil {
		t.Fatalf("Error parsing file: %v", err)
	}
	if formula != "sloc" {
		t.Fatalf("Expected formula sloc, got %s", formula)
	}
}

func TestParseFileSphericalLawOfCosines(t *testing.T) {
	filePath, err := makeTestFile(validFileSphericalLawOfCosines)
	if err != nil {
		t.Fatalf("Error creating test file: %v", err)
	}
	defer os.Remove(filePath)

	// read the test data from the file
	_, _, _, _, formula, err := ParseFile(filePath)
	if err != nil {
		t.Fatalf("Error parsing file: %v", err)
	}
	if formula != "sloc" {
		t.Fatalf("Expected formula sloc, got %s", formula)
	}
}

func TestParseFileHaversine(t *testing.T) {
	filePath, err := makeTestFile(validFileHaversine)
	if err != nil {
		t.Fatalf("Error creating test file: %v", err)
	}
	defer os.Remove(filePath)

	// read the test data from the file
	_, _, _, _, formula, err := ParseFile(filePath)
	if err != nil {
		t.Fatalf("Error parsing file: %v", err)
	}
	if formula != "haversine" {
		t.Fatalf("Expected formula haversine, got %s", formula)
	}
}

func TestParseFileEmptyFormula(t *testing.T) {
	filePath, err := makeTestFile(emptyFormulaFile)
	if err != nil {
		t.Fatalf("Error creating test file: %v", err)
	}
	defer os.Remove(filePath)

	// read the test data from the file
	_, _, _, _, formula, err := ParseFile(filePath)
	if err != nil {
		t.Fatalf("Error parsing file: %v", err)
	}
	if formula != "vincenty" {
		t.Fatalf("Expected formula vincenty, got %s", formula)
	}
}

func TestParseFileNoFile(t *testing.T) {
	filePath := "nonexistent.json"
	_, _, _, _, _, err := ParseFile(filePath)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestParseFileInvalidJSON(t *testing.T) {
	filePath, err := makeTestFile(invalidJSONFile)
	if err != nil {
		t.Fatalf("Error creating test file: %v", err)
	}
	defer os.Remove(filePath)

	// read the test data from the file
	_, _, _, _, _, err = ParseFile(filePath)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestParseFileInvalidLatitude(t *testing.T) {
	filePath, err := makeTestFile(invalidLatitudeFile)
	if err != nil {
		t.Fatalf("Error creating test file: %v", err)
	}
	defer os.Remove(filePath)

	// read the test data from the file
	_, _, _, _, _, err = ParseFile(filePath)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestParseFileInvalidLongitude(t *testing.T) {
	filePath, err := makeTestFile(invalidLongitudeFile)
	if err != nil {
		t.Fatalf("Error creating test file: %v", err)
	}
	defer os.Remove(filePath)

	// read the test data from the file
	_, _, _, _, _, err = ParseFile(filePath)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}
