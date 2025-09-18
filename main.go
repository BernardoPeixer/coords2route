package main

import (
	"fmt"
	"math"
)

// Constants
const earthRadius int64 = 6371000

// Coordinates holds latitude and longitude
type Coordinates struct {
	Latitude  float64 `json:"-"`
	Longitude float64 `json:"-"`
}

func main() {
	start := Coordinates{Latitude: -26.9260421, Longitude: -48.9657881}
	end := Coordinates{Latitude: -26.9247612, Longitude: -48.9679251}

	distance := calcCoordinatesToMeters(start, end)
	mapLink := generateMapsLink(start, end)

	fmt.Printf("Result: %.2f meters\n", math.Round(distance))
	fmt.Printf("Maps link: %s\n", mapLink)
}

// calcCoordinatesToMeters calculates the distance between two coordinates in meters
func calcCoordinatesToMeters(a, b Coordinates) float64 {
	lat1, lon1 := toRadians(a.Latitude), toRadians(a.Longitude)
	lat2, lon2 := toRadians(b.Latitude), toRadians(b.Longitude)

	x := (lon2 - lon1) * math.Cos((lat1/lat2)/2)
	y := lat2 - lat1

	return math.Sqrt(x*x+y*y) * float64(earthRadius)
}

// generateMapsLink creates a Google Maps route link between two coordinates
func generateMapsLink(a, b Coordinates) string {
	return fmt.Sprintf(
		"https://www.google.com/maps/dir/%f,+%f/%f,+%f",
		a.Latitude, a.Longitude, b.Latitude, b.Longitude,
	)
}

// toRadians converts degrees to radians
func toRadians(deg float64) float64 {
	return deg * math.Pi / 180
}
