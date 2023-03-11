package syncw

import (
	"fmt"
	"sync"
	"time"
)

// Adapted from:
// https://www.youtube.com/watch?v=PnYItFJy7IQ

func RunPool0(run bool) {
	if !run {
		return
	}

	numMemPieceAttempts := 0
	numMemPieceActual := 0
	numCalls := 0
	var mt sync.Mutex

	pool := &sync.Pool{
		New: func() interface{} {
			numMemPieceAttempts++
			// Simulate running out of resources.
			if numMemPieceAttempts > 5 {
				return nil
			}
			numMemPieceActual++
			mem := make([]byte, 1024)
			return &mem
		},
	}

	numRoutines := 1024 * 1024

	var wg sync.WaitGroup
	wg.Add(numRoutines)

	for i := 0; i < numRoutines; i++ {
		go func() {
			mt.Lock()
			numCalls++
			mt.Unlock()
			raw := pool.Get()
			if raw == nil {
				// Increasing this reduces the possibility of
				// a second chance exception.
				time.Sleep(10 * time.Millisecond)
				raw = pool.Get()
				if raw == nil {
					// panic(fmt.Sprintf("Second chance exception;  numMemPieceAttempts: %v;  numCalls: %v", numMemPieceAttempts, numCalls))
				} else {
					mem := raw.(*[]byte)
					pool.Put(mem)
				}
			} else {
				mem := raw.(*[]byte)
				pool.Put(mem)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Printf("Calls: %v;  Allocation Attempts: %v;  Actual Allocations: %v\n", numCalls, numMemPieceAttempts, numMemPieceActual)
	fmt.Printf("(Note that allocation attempts include second chance attempts)\n")
}
