package main

import (
	"fmt"
	"math/rand"
)

func thirder() (int, int, int) {
	if rand.Intn(10) == 1 {
		return 0, 1, 0
	}
	return 0, 1, 1
}

func halfer() (int, int, int) {
	if rand.Intn(10) == 1 {
		return 0, 1, 0
	}
	return 0, 0, 1
}

func trials(
	trial func() (int, int, int),
	count int,
) (int, int, int) {
	var s, m, t int
	for i := 0; i < count; i++ {
		s0, m0, t0 := trial()
		s += s0
		m += m0
		t += t0
	}
	return s, m, t
}

func main3() {

	count := 1000000

	s, m, t := trials(thirder, count)
	fmt.Printf("%v %v %v => %v %v %v\n", s, m, t, s+t, s+t+m, float64(s+t)/float64(s+t+m))

	s, m, t = trials(halfer, count)
	fmt.Printf("%v %v %v => %v %v %v\n", s, m, t, s+t, s+t+m, float64(s+t)/float64(s+t+m))
}
