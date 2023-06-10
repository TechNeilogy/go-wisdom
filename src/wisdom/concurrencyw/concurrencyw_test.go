package concurrencyw

import (
	"fmt"
	"time"
)

func Example_basic() {

	start := time.Now()

	done := make(chan bool)
	go Basic(done, true, 84)
	<-done

	fmt.Printf(" [%v]", time.Since(start))

	start = time.Now()

	go Basic(done, false, 84)
	<-done

	fmt.Printf(" [%v]", time.Since(start))

	//output: ?
}
