package concurrencyw

import (
	"fmt"
	"sync"
	"time"
)

func ExampleBasicChan() {

	start := time.Now()

	done := make(chan bool)
	go BasicChan(done, true, 84)
	<-done

	fmt.Printf(" [%v]", time.Since(start))

	start = time.Now()

	go BasicChan(done, false, 84)
	<-done

	fmt.Printf(" [%v]", time.Since(start))

	//output: ?
}

func ExampleBasicWaitGroup() {

	start := time.Now()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go BasicWaitGroup(&wg, true, 84)
	wg.Wait()

	fmt.Printf(" [%v]", time.Since(start))

	start = time.Now()

	wg.Add(1)
	go BasicWaitGroup(&wg, false, 84)
	wg.Wait()

	fmt.Printf(" [%v]", time.Since(start))

	//output: ?
}
