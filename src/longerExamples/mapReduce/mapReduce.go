package mapReduce

import (
	"context"
	"sync"
)

// MapFunc is a map function: T => U
type MapFunc[T any, U any] func(
	context.Context,
	T,
) U

// Iter defines an iterable.
// Go deliberately lacks one as a general case,
// but they can sometimes be useful.
type Iter[T any] interface {
	More() bool
	Item() T
}

// ReduceFunc is a reduce function: U, V => V
type ReduceFunc[U any, V any] func(
	context.Context,
	U,
	V,
) V

// reduce wraps a ReduceFunc in channel logic.
func (rf ReduceFunc[U, V]) reduce(
	ctx context.Context,
	cU chan U,
	cQ chan bool,
	cV chan V,
) {
	var acc V
	for {
		select {
		case x := <-cU:
			// call reduce function
			acc = rf(ctx, x, acc)
		case <-cQ:
			// quit
			cV <- acc
			return
		}
	}
}

// MapReduce runs a parallel map/reduce using goroutines.
//
// mf is the mapping function,
// rf is the reduce function, and
// iter is the iterator for the mapped items.
func MapReduce[T any, U any, V any](
	ctx context.Context,
	mf MapFunc[T, U],
	rf ReduceFunc[U, V],
	iter Iter[T],
) V {
	cU := make(chan U)    // map => reduce
	cQ := make(chan bool) // signal to quit reduce goroutine
	cV := make(chan V)    // reduce => this function

	// Spin up the reduce goroutine.
	go rf.reduce(ctx, cU, cQ, cV)

	// detect all map goroutines have completed
	var wg sync.WaitGroup

	// spin up the map goroutines
	for iter.More() {
		wg.Add(1)
		item := iter.Item() // prevents a race condition
		go func() {
			defer wg.Done()
			cU <- mf(ctx, item)
		}()
	}

	wg.Wait()

	// signal the reduce goroutine to stop
	cQ <- true

	// return the result of the reduce goroutine
	return <-cV
}

// IterList is an iterator over a list.
type IterList[V any] struct {
	l []V
	i int
}

// MakeIterList creates an iterator over a map.
func MakeIterList[V any](
	l []V,
) IterList[V] {
	return IterList[V]{l, -1}
}

// IterList More fulfills the Iter interface.
func (il *IterList[V]) More() bool {
	if il.i < len(il.l)-1 {
		il.i++
		return true
	}
	return false
}

// IterList Item fulfills the Iter interface.
func (il *IterList[V]) Item() V {
	if il.i < len(il.l) {
		return il.l[il.i]
	}
	var v V
	return v
}

// IterMap is an iterator over a map.
type IterMap[K comparable, V any] struct {
	m map[K]V
	k []K
	i int
}

// MakeIterMap creates an iterator over a map.
func MakeIterMap[K comparable, V any](
	m map[K]V,
) IterMap[K, V] {
	var ks []K
	for k := range m {
		ks = append(ks, k)
	}
	return IterMap[K, V]{m, ks, -1}
}

// IterMap More fulfills the Iter interface.
func (im *IterMap[K, V]) More() bool {
	if im.i < len(im.k)-1 {
		im.i++
		return true
	}
	return false
}

// IterMap Item fulfills the Iter interface.
func (im *IterMap[K, V]) Item() V {
	if im.i < len(im.k) {
		return im.m[im.k[im.i]]
	}
	var v V
	return v
}
