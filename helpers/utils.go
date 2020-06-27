package helpers

import "math"

// GetWeightedValue lol
func GetWeightedValue(vals []float64, decay float64) int {
	weightDecay := 1.0
	result := 0.00
	for i := 0; i < len(vals); i++ {
		result += weightDecay * vals[i]
		weightDecay *= decay
	}
	return int(math.Round(result * (1.0 - decay)))
}
