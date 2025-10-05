package main

import (
	"math"
	"math/rand/v2"
)

func getRandomPoint(dist Distribution, radius float64) (point Point) {
	theta := rand.Float64() * 2 * math.Pi

	var rad float64
	if dist == Area {
		u := rand.Float64()
		rad = radius * math.Sqrt(u)
	} else {
		rad = rand.Float64() * radius
	}

	x := rad * math.Cos(theta)
	y := rad * math.Sin(theta)

	return Point{x, y}
}

func getDistance(point Point) float64 {
	return math.Sqrt(math.Pow(point.x, 2) + math.Pow(point.y, 2))
}

func getMean(subscribers []Subscriber) float64 {
	var mean float64

	for i := range subscribers {
		mean += subscribers[i].distance
	}

	return mean / float64(len(subscribers))
}

func getStdDeviation(subscribers []Subscriber) float64 {
	mean := getMean(subscribers)

	var sumSqDiff float64
	for _, value := range subscribers {
		diff := value.distance - mean
		sumSqDiff += math.Pow(diff, 2)
	}

	variance := sumSqDiff / float64(len(subscribers)-1)

	return math.Sqrt(variance)
}
