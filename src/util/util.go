package util

import (
	"fmt"
	"math"
	"strings"
)

func PrintHeader(msg string) {
	fmt.Println("")
	fmt.Printf("+%v+\n", strings.Repeat("-", len(msg)+2))
	fmt.Printf("| %v |\n", msg)
	fmt.Printf("+%v+\n", strings.Repeat("-", len(msg)+2))
	fmt.Println("")
}

func Stats(fs []float64) (float64, float64) {

	count := len(fs)

	if count < 1 {
		return 0, 0
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

	return mean, math.Sqrt(std)
}
