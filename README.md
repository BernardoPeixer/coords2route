# map-your-distance

Snap distances between points and get a Google Maps route instantly!

A simple Go program to calculate the distance between two geographic coordinates and generate a Google Maps directions link.

---

## Features

- Calculate distance between two coordinates in meters.
- Generate a Google Maps link with the route between points.
- Lightweight and easy to use.

---

## Example

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	firstCord := Coordinates{
		Latitude:  -26.9260421,
		Longitude: -48.9657881,
	}

	secondCord := Coordinates{
		Latitude:  -26.9247612,
		Longitude: -48.9679251,
	}

	result := calcCoordinatesToMeters(firstCord, secondCord)
	link := generateMapsLink(firstCord, secondCord)

	fmt.Printf("Result: %.2f meters\n", math.Round(result))
	fmt.Printf("Maps link: %s\n", link)
}
```

## Output

```
Result: 271 meters
Maps link: https://www.google.com/maps/dir/-26.9260421,+-48.9657881/-26.9247612,+-48.9679251
```

## Installation

- Make sure you have Go installed
- Clone the repository:

```
git clone https://github.com/BernardoPeixer/coords2route.git
```

- Run the program

```
go run main.go
```

## Functions

- calcCoordinatesToMeters(firstCoordinate, secondCoordinate Coordinates) float64
Calculates the distance in meters between two coordinates.

- generateMapsLink(firstCoordinate, secondCoordinate Coordinates) string Generates a Google Maps route link between two coordinates.

