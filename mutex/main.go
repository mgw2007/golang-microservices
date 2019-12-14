package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	counter       = 0
	lock          sync.Mutex
	atomicCounter = AtomicInt{}
	aa            = AtomicCounter{
		counter: 0,
	}
)

type Counter interface {
	Inc()
	Load() int64
}

// Atomic Implementation

type AtomicCounter struct {
	counter int64
}

func (c *AtomicCounter) Inc() {
	atomic.AddInt64(&c.counter, 1)
}

func (c *AtomicCounter) Load() int64 {
	return atomic.LoadInt64(&c.counter)
}

type AtomicInt struct {
	value int
	lock  sync.Mutex
}

func (i *AtomicInt) Increase() {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.value++
}

func (i *AtomicInt) Decrease() {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.value--
}

func (i *AtomicInt) Value() int {
	i.lock.Lock()
	defer i.lock.Unlock()
	return i.value
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go updateCounter(&wg)
	}
	wg.Wait()
	fmt.Println(fmt.Sprintf("final counter: %d", counter))
	fmt.Println(fmt.Sprintf("final atomic counter value: %d", atomicCounter.Value()))
	fmt.Println(fmt.Sprintf("Atomic counter value: %d", aa.Load()))
}

func updateCounter(wg *sync.WaitGroup) {
	lock.Lock()
	defer lock.Unlock()

	counter++

	atomicCounter.Increase()
	aa.Inc()
	wg.Done()
}
