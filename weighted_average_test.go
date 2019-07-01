package main

import (
	"math"
	"testing"
)

func TestWeightedAverage(t *testing.T) {
	entriesSet := []struct {
		values  []float64
		weights []float64
		exp     float64
	}{
		{
			[]float64{10, 11, 12},
			[]float64{0.5, 1, 2},
			11.4286,
		},
	}

	for _, entries := range entriesSet {
		result := WeightedArithmeticMean(entries.values, entries.weights)
		if !areEqual(entries.exp, result) {
			t.Errorf("Expected %f; got %f", entries.exp, result)
		}
	}
}

const tolerance = 0.0001

func areEqual(a float64, b float64) bool {
	diff := math.Abs(a - b)

	return diff < tolerance
}
