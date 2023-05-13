package panicw

import "fmt"

func a(doPanic bool) {
	fmt.Println("a0")
	if doPanic {
		panic("a")
	}
	fmt.Println("a1")
}

func b(doPanic, doRecover bool) {
	local := fmt.Sprintf("%v %v", doPanic, doRecover)
	if doRecover {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("br0", local)
			} else {
				fmt.Println("br1", local)
			}
		}()
	}
	fmt.Println("b0")
	a(doPanic)
	fmt.Println("b1")
}

func Test() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("mr0")
		} else {
			fmt.Println("mr1")
		}
	}()

	fmt.Println("m0")
	b(false, false)
	fmt.Println("m1")
	b(false, true)
	fmt.Println("m2")
	b(true, true)
	fmt.Println("m3")
	b(true, false)
	fmt.Println("m4")
}
