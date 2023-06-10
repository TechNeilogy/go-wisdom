package concurrencyw

import (
	"fmt"
	"sync"
)

func Fib(n int) int {
	if n > 42 {
		panic(fmt.Sprintf("did you really mean to compute fib(%v)?", n))
	}
	if n < 2 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}

func RunFib(
	wg *sync.WaitGroup,
	count int,
) {
	// If you want to check order.
	//fmt.Printf(" (%v)", count)
	//fmt.Printf(" %v", Fib(count/2))
	Fib(count / 2)
	wg.Done()
}

func Basic(
	done chan bool,
	concurrent bool,
	count int,
) {
	wg := sync.WaitGroup{}
	for count > 0 {
		wg.Add(1)
		if concurrent {
			go RunFib(&wg, count)
		} else {
			RunFib(&wg, count)
		}
		count--
	}
	wg.Wait()

	done <- true
}
