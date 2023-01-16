package structw

import (
	"fmt"
	"github.com/TechNeilogy/go-wisdom/src/util"
	"time"
)

type intStruct struct {
	a int
}

type stringStruct struct {
	b string
}

func testIntStruct(count int) (rtn []intStruct) {
	for i := 0; i < count; i++ {
		rtn = append(rtn, intStruct{i})
	}
	return
}

func testStringStruct(count int, text string) (rtn []stringStruct) {
	for i := 0; i < count; i++ {
		rtn = append(rtn, stringStruct{text})
	}
	return
}

func RunStringStructPenalty(run bool) {
	if !run {
		return
	}

	util.PrintHeader("String struct Performance Penalty")

	runs := 5
	count := 1000
	countInner := 100
	countTotal := count * count * countInner
	text := ""

	intRun := []float64{}
	stringRun := []float64{}

	fmt.Printf("Outer Loops: %v\n", count*count)
	fmt.Printf("Inner Loops: %v\n", countInner)
	fmt.Printf("Total Loops: %v\n", countTotal)
	fmt.Printf("Run: ")
	for k := 1; k <= runs; k++ {

		fmt.Printf(" %v", k)

		sum := 0
		start := time.Now()
		for i := 0; i < count; i++ {
			for j := 0; j < count; j++ {
				sum += len(testIntStruct(countInner))
			}
		}
		intRun = append(intRun, time.Since(start).Seconds())

		sum = 0
		start = time.Now()
		for i := 0; i < count; i++ {
			for j := 0; j < count; j++ {
				sum += len(testStringStruct(countInner, text))
			}
		}
		stringRun = append(stringRun, time.Since(start).Seconds())
	}
	fmt.Printf("\n")

	fmt.Printf("Result of %v Runs:\n", runs)
	sumInt, meanInt, stdInt := util.Stats(intRun)
	sumString, meanString, stdString := util.Stats(stringRun)
	fmt.Printf("    intStruct:  Sum=%v Mean=%v  Std=%v\n", sumInt, meanInt, stdInt)
	fmt.Printf("    stringStruct:  Sum=%v Mean=%v  Std=%v\n", meanString, stdString, sumString)
	fmt.Printf("    intStruct per Second:  %v\n", float64(countTotal)/sumInt)
	fmt.Printf("    stringStruct per Second:  %v\n", float64(countTotal)/sumString)

}
