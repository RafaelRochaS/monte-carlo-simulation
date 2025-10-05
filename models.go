package main

type Subscriber struct {
	location Point
	distance float64
}

type Point struct {
	x float64
	y float64
}

type Distribution int

const (
	Area Distribution = iota
	Distance
)

func (d Distribution) String() string {
	switch d {
	case Area:
		return "Area"
	case Distance:
		return "Distance"
	default:
		return "Unknown"
	}
}
