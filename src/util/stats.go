package util

import (
	"math"
)

// Stats computes the count, sum, mean, and standard deviation of data.
func Stats(data []float64) (float64, float64, float64, float64) {

	count := len(data)

	if count < 1 {
		return 0, 0, 0, 0
	}

	var sum float64
	for _, f := range data {
		sum += f
	}

	countf := float64(count)

	mean := sum / countf

	var std float64
	for _, f := range data {
		d := f - mean
		std += d * d
	}

	return countf, sum, mean, math.Sqrt(std)
}
