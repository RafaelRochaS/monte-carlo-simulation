package main

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"image/color"
	"math"
	"math/rand"
	"sort"
)

type Subscriber struct {
	location Point
	distance float64
}

type Point struct {
	x float64
	y float64
}

const Radius = 2

var totalSubscribers = math.Pow10(Radius)

func main() {

	subscribers := make([]Subscriber, int(totalSubscribers))

	for i := range subscribers {
		location := getRandomPoint()
		subscribers[i] = Subscriber{
			location: location,
			distance: getDistance(location),
		}
	}

	sort.Slice(subscribers, func(i, j int) bool {
		return subscribers[i].distance < subscribers[j].distance
	})

	printSubscribers(subscribers)
	fmt.Println("Total subscribers:", int(totalSubscribers))
	fmt.Printf("\nMean: %.3f\n", getMean(subscribers))
	fmt.Printf("\nStandard deviation: %.3f\n", getStdDeviation(subscribers))
	err := plotDistribution(subscribers)

	if err != nil {
		fmt.Printf("Failed to plot distribution: %v", err)
	}

	err = plotSubscribers(subscribers)

	if err != nil {
		fmt.Printf("Failed to plot subscribers: %v", err)
	}
}

func getRandomPoint() (point Point) {
	theta := rand.Float64() * 2 * math.Pi
	u := rand.Float64()
	rad := Radius * math.Sqrt(u)
	x := rad * math.Cos(theta)
	y := rad * math.Sin(theta)

	return Point{x, y}

}

func getDistance(point Point) float64 {
	return math.Sqrt(math.Pow(point.x, 2) + math.Pow(point.y, 2))
}

func printSubscribers(subscribers []Subscriber) {
	for i := range subscribers {
		fmt.Printf("Subscriber: %+v\n", subscribers[i])
	}
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

func plotDistribution(subscribers []Subscriber) error {
	p := plot.New()
	p.Title.Text = "Subscribers Distribution"

	v := make(plotter.Values, len(subscribers))
	for i := range v {
		v[i] = subscribers[i].distance
	}

	h, err := plotter.NewHist(v, 16)
	if err != nil {
		return err
	}

	h.Normalize(1)
	p.Add(h)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "distribution.png"); err != nil {
		return err
	}

	return nil
}

func plotSubscribers(subscribers []Subscriber) error {
	p := plot.New()

	pts := make(plotter.XYs, len(subscribers))

	for i := range pts {
		pts[i].X = subscribers[i].location.x
		pts[i].Y = subscribers[i].location.y
	}

	p.Title.Text = "Subscribers"
	p.X.Label.Text = "X-axis"
	p.Y.Label.Text = "Y-axis"

	s, err := plotter.NewScatter(pts)
	if err != nil {
		return err
	}

	s.Color = color.RGBA{R: 255, A: 255}
	s.Shape = draw.CircleGlyph{}
	s.Radius = vg.Points(.8)

	p.Add(s)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "subscribers.png"); err != nil {
		return err
	}

	return nil
}
