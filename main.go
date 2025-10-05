package main

import (
	"fmt"
	"math"
	"sort"
)

const Radius = 2

var totalSubscribers = math.Pow10(Radius)

func main() {

	subscribersArea := makeSlice(Area)
	subscribersDistance := makeSlice(Distance)

	fmt.Println("Total subscribers:", int(totalSubscribers))

	fmt.Printf("\nMean - Area: %.3f\n", getMean(subscribersArea))
	fmt.Printf("Mean - Distance: %.3f\n", getMean(subscribersDistance))

	fmt.Printf("Standard deviation - Area: %.3f\n", getStdDeviation(subscribersArea))
	fmt.Printf("Standard deviation - Distance: %.3f\n", getStdDeviation(subscribersDistance))

	err := plotDistribution(subscribersArea, Area, Radius)
	err = plotDistribution(subscribersDistance, Distance, Radius)

	if err != nil {
		fmt.Printf("Failed to plot distribution: %v", err)
	}

	err = plotSubscribers(subscribersArea, Area, Radius)
	err = plotSubscribers(subscribersDistance, Distance, Radius)

	if err != nil {
		fmt.Printf("Failed to plot subscribers: %v", err)
	}
}

func makeSlice(dist Distribution) (slice []Subscriber) {
	for i := range slice {
		location := getRandomPoint(dist)
		slice[i] = Subscriber{
			location: location,
			distance: getDistance(location),
		}
	}

	sort.Slice(slice, func(i, j int) bool {
		return slice[i].distance < slice[j].distance
	})

	return
}
