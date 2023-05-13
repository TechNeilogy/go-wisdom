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

type intPtrStruct struct {
	b *int
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

func testIntPtrStruct(count int, datum *int) (rtn []intPtrStruct) {
	for i := 0; i < count; i++ {
		rtn = append(rtn, intPtrStruct{datum})
	}
	return
}

// StringStructPenalty shows the high performance penalty
// from passing structures with strings.
func StringStructPenalty() {

	runs := 5
	count := 100
	countInner := 10
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
	sumInt, meanInt, _ := util.Stats(intRun)
	sumString, meanString, _ := util.Stats(stringRun)
	fmt.Printf("sumString > sumInt = %v\n", sumString > sumInt)
	fmt.Printf("meanString > meanInt = %v", meanString > meanInt)

	// Can't be tested, because they vary,
	// but if you want to see the results:
	//
	//fmt.Printf("\n    intStruct:  Sum=%v Mean=%v\n", sumInt, meanInt)
	//fmt.Printf("    stringStruct:  Sum=%v Mean=%v\n", meanString, sumString)
	//fmt.Printf("    intStruct per Second:  %v\n", float64(countTotal)/sumInt)
	//fmt.Printf("    stringStruct per Second:  %v\n", float64(countTotal)/sumString)
}

// IntPtrStructPenalty shows low the performance penalty
// from passing structures with pointers.
func IntPtrStructPenalty() {

	runs := 5
	count := 100
	countInner := 10
	countTotal := count * count * countInner

	intRun := []float64{}
	intPtrRun := []float64{}

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

		datum := 22
		sum = 0
		start = time.Now()
		for i := 0; i < count; i++ {
			for j := 0; j < count; j++ {
				sum += len(testIntPtrStruct(countInner, &datum))
			}
		}
		intPtrRun = append(intPtrRun, time.Since(start).Seconds())
	}
	fmt.Printf("\n")

	fmt.Printf("Result of %v Runs:\n", runs)
	sumInt, meanInt, _ := util.Stats(intRun)
	sumIntPtr, meanIntPtr, _ := util.Stats(intPtrRun)
	fmt.Printf("sumIntPtr > sumInt = %v\n", sumIntPtr > sumInt)
	fmt.Printf("meanIntPtr > meanInt = %v\n", meanIntPtr > meanInt)

	// Can't be tested, because they vary,
	// but if you want to see the results:
	//
	//fmt.Printf("\n    intStruct:  Sum=%v Mean=%v\n", sumInt, meanInt)
	//fmt.Printf("    IntPtrStruct:  Sum=%v Mean=%v\n", sumIntPtr, meanIntPtr)
	//fmt.Printf("    intStruct per Second:  %v\n", float64(countTotal)/sumInt)
	//fmt.Printf("    IntPtrStruct per Second:  %v\n", float64(countTotal)/sumIntPtr)
}
