package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	buff := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)

	go sender(buff)
	go receiver(buff, &wg)

	fmt.Println("waiting...")
	wg.Wait()
}

func receiver(buff chan int, wg *sync.WaitGroup) {
	for value := range buff {
		fmt.Println("receiver:", value)
		time.Sleep(time.Second)
	}
	fmt.Println("receiver-1")
	wg.Done()
	fmt.Println("receiver-2")
	return
}

func sender(buff chan int) {
	for i := 0; i < 3; i++ {
		buff <- i
		fmt.Println("sender:", i)
	}
	fmt.Println("sender-1")
	close(buff)
	fmt.Println("sender-2")
	return
}
