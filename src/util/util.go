package util

import (
	"math"
)

func Stats(fs []float64) (float64, float64, float64) {

	count := len(fs)

	if count < 1 {
		return 0, 0, 0
	}

	var sum float64
	for _, f := range fs {
		sum += f
	}

	mean := sum / float64(count)

	var std float64
	for _, f := range fs {
		d := f - mean
		std += d * d
	}

	return sum, mean, math.Sqrt(std)
}
