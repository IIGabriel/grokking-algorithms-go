package chapter8

import (
	"math"
)

type City struct {
	X, Y float64
}

// distanceTo calculates the Euclidean distance between two cities
func (s *City) distanceTo(other City) float64 {
	dx := s.X - other.X
	dy := s.Y - other.Y
	return math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
}

// TravelingSalesmanProblem finds an approximate solution to the Traveling Salesman Problem using a greedy algorithm.
// Starts from the first city, and at each step, visits the nearest unvisited city.
// Takes a slice of City structs representing the cities to be visited.
// Returns the total distance of the approximate shortest path that visits each city exactly once and returns to the starting city.
// - Time complexity: O(n^2)
// - Space complexity: O(n)
func TravelingSalesmanProblem(cities []City) float64 {
	if len(cities) == 0 {
		return 0
	}
	var totalDistance float64
	visited := make([]bool, len(cities))
	currentCity := cities[0]
	visited[0] = true

	for range len(cities) - 1 {
		indexLowerDistance := -1
		lowerDistance := math.MaxFloat64
		for j, subCity := range cities {
			if visited[j] {
				continue
			}
			if distance := currentCity.distanceTo(subCity); lowerDistance > distance {
				lowerDistance = distance
				indexLowerDistance = j
			}
		}
		visited[indexLowerDistance] = true
		totalDistance += lowerDistance
		currentCity = cities[indexLowerDistance]
	}
	return totalDistance
}
