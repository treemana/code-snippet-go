package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var test int64
	var wg sync.WaitGroup

	for j := 0; j < 100; j++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10000; i++ {
				atomic.AddInt64(&test, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(test)
}
