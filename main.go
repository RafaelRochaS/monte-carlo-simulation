package main

import "fmt"

const MaxRadius = 6

func main() {
	for radius := 1; radius <= MaxRadius; radius++ {
		err := simulate(radius)

		if err != nil {
			fmt.Printf("Simulation failed for radius %d: %v", radius, err)
		}
	}
}
