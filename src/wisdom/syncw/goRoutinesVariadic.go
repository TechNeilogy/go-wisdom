package syncw

import (
	"fmt"
	"sync"
)

// Wisdom adapted from:
// https://blogtitle.github.io/go-advanced-concurrency-patterns-part-1/

func firstComeFirstServedGoroutinesVariadic(
	message interface{},
	chans []chan interface{},
) {
	var wg sync.WaitGroup
	wg.Add(len(chans))
	for i, c := range chans {
		c := c
		i := i
		go func() {
			c <- message
			fmt.Printf(" %v", i)
			wg.Done()
		}()
	}
	wg.Wait()
}

func RunFirstComeFirstServedGoroutinesVariadic(run bool) {
	if !run {
		return
	}

	size := 10
	chans := make([]chan interface{}, size)
	for i := range chans {
		chans[i] = make(chan interface{}, 1)
	}

	fmt.Printf("Go Routines:")
	firstComeFirstServedGoroutinesVariadic(
		"",
		chans)
	fmt.Printf("\n")
}
