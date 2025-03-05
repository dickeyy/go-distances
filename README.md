# go-distances

[![Go Reference](https://pkg.go.dev/badge/github.com/dickeyy/go-distances.svg)](https://pkg.go.dev/github.com/dickeyy/go-distances)
[![codecov](https://codecov.io/gh/dickeyy/go-distances/graph/badge.svg?token=6N8HYISOL0)](https://codecov.io/gh/dickeyy/go-distances)

A simple Go program to calculate circular distances between geographical points using different formulas.

## Description

`go-distances` calculates the distances between a set of geographical points arranged in a circular path. It prompts the user for the number of points, their latitudes and longitudes, the Earth's radius, and the formula to use for distance calculation. Currently, it supports the Haversine and Vincenty formulas.

## Features

- Calculates circular distances for any number of points.
- Supports Haversine and Vicinity distance formulas.
- User-friendly command-line interface.

## Prerequisites

- Go (version 1.16 or higher)

## Installation

Either, clone this repo, or download the binary.

### Clone the repo

```
git clone https://github.com/dickeyy/go-distances.git
cd go-distances
```

### Download the binary

```
go install github.com/dickeyy/go-distances@latest
```

## Usage

1.  **Run the program:**

    If you cloned the repo, run `go run main.go`.
    If you downloaded the binary, run `go-distances`.

2.  **Follow the prompts:**

    - Do you want to import points from a file? (y/n): y
    - Enter the path to the file.

    OR

    - Do you want to import points from a file? (y/n): n
    - Enter the number of points.
    - Enter the latitude and longitude for each point.
    - Enter the Earth's radius (in your desired unit, e.g., 6371 for kilometers).
    - Enter the formula to use (`haversine` or `vincenty`).

3.  **View the results:**

    The program will output the calculated distances between each pair of consecutive points in the circular path.

### Importing data from a file

If you wish to import data from a file, the file MUST be in JSON format, and follow the example format below:

```json
{
    "places": [
        {
            "name": "some name",
            "latitude": "some latitude", // must be a string in degrees
            "longitude": "some longitude"
        },
        {
            "name": "some other name",
            "latitude": "some other latitude",
            "longitude": "some other longitude"
        }
        ...
    ],
    "earthRadius": 12345, // whatever unit you want
    "formula": "haversine" // optional, defaults to "vincenty"
}
```

Note the comments, the coordinates must be in degrees and represented as strings. The `earthRadius` is the radius of the Earth in whatever unit you want, and the `formula` is optional and defaults to `vincenty`.

### test-all-data.sh

This script will run the program on all the test data files in the `test-data` directory, if you have a bunch of test data files you want to run quickly. Note: for this to work, you need to build the Go program first. To build the program, run `go build -o go-distances main.go`.

## Formulas

More indepth information about the formulas can be found [Here](./formulas/README.md).

- `haversine`: Uses the Haversine formula to calculate the distance between two points on a sphere.
- `vincenty`: Uses the Vincenty formula to calculate the distance between two points on a sphere.
- `sloc`: Uses the Spherical Law of Cosines formula to calculate the distance between two points on a sphere.

## License

MIT License, see [LISENCE file](./LICENSE).
