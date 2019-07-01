package main

// WeightedArithmeticMean of an array of values supplied with an array
// of weights. If there are more weights than values, the superfluous weights
// are discarded; if there are more values than weights, the superfluous values
// are discarded.
func WeightedArithmeticMean(values []float64, weights []float64) float64 {
	// Find shortest array
	var n int
	if len(values) < len(weights) {
		n = len(values)
	} else {
		n = len(weights)
	}

	weightedValueSum := float64(0)
	weightSum := float64(0)
	for i := 0; i < n; i++ {
		weightedValueSum += values[i] * weights[i]
		weightSum += weights[i]
	}

	result := weightedValueSum / weightSum

	return result
}
