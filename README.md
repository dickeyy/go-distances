# go-distances

A simple Go program to calculate circular distances between geographical points using different formulas.

## Description

`go-distances` calculates the distances between a set of geographical points arranged in a circular path. It prompts the user for the number of points, their latitudes and longitudes, the Earth's radius, and the formula to use for distance calculation. Currently, it supports the Haversine and Vicinity formulas.

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

    - Enter the number of points.
    - Enter the latitude and longitude for each point.
    - Enter the Earth's radius (in your desired unit, e.g., 6371 for kilometers).
    - Enter the formula to use (`haversine` or `vicinity`).

3.  **View the results:**

    The program will output the calculated distances between each pair of consecutive points in the circular path.

## Formulas

- **Haversine:** Calculates the great-circle distance between two points on a sphere given their longitudes and latitudes. This is a common and accurate method for calculating distances on Earth.
- **Vicinity:** (Describe the vicinity formula if you have details about it. Otherwise, state that it's an alternative distance calculation method.)

## License

MIT License, see [LISENCE file](./LICENSE).
