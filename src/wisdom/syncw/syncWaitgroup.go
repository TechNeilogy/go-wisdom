package syncw

import (
	"fmt"
	"sync"
)

var count int

func gr(
	wg *sync.WaitGroup,
	mt *sync.Mutex,
) {
	defer wg.Done()
	mt.Lock()
	count += 1
	mt.Unlock()
}

func syncWaitgroup0() {

	var wg sync.WaitGroup
	var mt sync.Mutex

	count = 0

	size := 1000

	for i := 0; i < size; i++ {
		wg.Add(1)
		go gr(&wg, &mt)
	}

	wg.Wait()

	fmt.Printf("Wait Group 0, expected: %v; got: %v\n", size, count)
}

func syncWaitgroup1() {

	var wg sync.WaitGroup
	var mt sync.Mutex

	count = 0

	size := 1000
	wg.Add(size)

	for i := 0; i < size; i++ {
		go gr(&wg, &mt)
	}

	wg.Wait()

	fmt.Printf("Wait Group 1, expected: %v; got: %v\n", size, count)
}

func RunSyncWaitGroup(run bool) {
	if !run {
		return
	}

	syncWaitgroup0()
	syncWaitgroup1()
}
