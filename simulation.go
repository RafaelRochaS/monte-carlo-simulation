package main

import (
	"fmt"
	"math"
	"sort"
)

func simulate(radius int) error {
	fmt.Println("Starting simulation for radius", radius, "...")
	totalSubscribers := int(math.Pow10(radius))

	subscribersArea := makeSlice(Area, totalSubscribers, radius)
	subscribersDistance := makeSlice(Distance, totalSubscribers, radius)

	fmt.Println("Total subscribers:", int(totalSubscribers))
	fmt.Printf("\nMean - Area: %.3f\n", getMean(subscribersArea))
	fmt.Printf("Mean - Distance: %.3f\n", getMean(subscribersDistance))

	fmt.Printf("Standard deviation - Area: %.3f\n", getStdDeviation(subscribersArea))
	fmt.Printf("Standard deviation - Distance: %.3f\n", getStdDeviation(subscribersDistance))

	err := plotDistribution(subscribersArea, Area, radius)
	err = plotDistribution(subscribersDistance, Distance, radius)

	if err != nil {
		fmt.Printf("Failed to plot distribution: %v", err)
	}

	err = plotSubscribers(subscribersArea, Area, radius)
	err = plotSubscribers(subscribersDistance, Distance, radius)

	if err != nil {
		fmt.Printf("Failed to plot subscribers: %v", err)
	}

	return nil
}

func makeSlice(dist Distribution, size int, radius int) (slice []Subscriber) {
	slice = make([]Subscriber, size)

	for i := range slice {
		location := getRandomPoint(dist, float64(radius))

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
