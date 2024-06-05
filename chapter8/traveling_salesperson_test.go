package chapter8

import (
	"testing"
)

func TestTravelingSalesmanProblem(t *testing.T) {
	type testModel struct {
		cities   []City
		expected float64
	}

	testCase := func(t *testing.T, testParams testModel) {
		got := TravelingSalesmanProblem(testParams.cities)
		if got != testParams.expected {
			t.Errorf("TravelingSalesmanProblem(%v) = %v; want %v", testParams.cities, got, testParams.expected)
		}
	}

	t.Run("no cities", func(t *testing.T) {
		testParam := testModel{
			cities:   []City{},
			expected: 0,
		}
		testCase(t, testParam)
	})

	t.Run("single city", func(t *testing.T) {
		testParam := testModel{
			cities:   []City{{X: 0, Y: 0}},
			expected: 0,
		}
		testCase(t, testParam)
	})

	t.Run("two cities", func(t *testing.T) {
		testParam := testModel{
			cities:   []City{{X: 0, Y: 0}, {X: 3, Y: 4}},
			expected: 5,
		}
		testCase(t, testParam)
	})

	t.Run("three cities forming a triangle", func(t *testing.T) {
		testParam := testModel{
			cities: []City{
				{X: 0, Y: 0},
				{X: 3, Y: 4},
				{X: 6, Y: 0},
			},
			expected: 10, // 5 + 5
		}
		testCase(t, testParam)
	})

	t.Run("four cities forming a square", func(t *testing.T) {
		testParam := testModel{
			cities: []City{
				{X: 0, Y: 0},
				{X: 0, Y: 4},
				{X: 4, Y: 0},
				{X: 4, Y: 4},
			},
			expected: 12, // 4 + 4 + 4 (visiting nearest unvisited city)
		}
		testCase(t, testParam)
	})

	t.Run("multiple cities in a straight line", func(t *testing.T) {
		testParam := testModel{
			cities: []City{
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: 2, Y: 0},
				{X: 3, Y: 0},
				{X: 4, Y: 0},
			},
			expected: 4, // 1 + 1 + 1 + 1 (visiting nearest unvisited city)
		}
		testCase(t, testParam)
	})
}
