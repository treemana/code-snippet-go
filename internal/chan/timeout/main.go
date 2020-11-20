package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	var rtvChan = make(chan int64, 1)
	go delayFun(rtvChan)
	select {
	case <-ctx.Done():
		fmt.Println("main")
	case rtv := <-rtvChan:
		fmt.Println(rtv)
	}
	fmt.Println("finish")
}

func delayFun(rtv chan int64) {
	time.Sleep(time.Second * 4)
	rtv <- 3
	fmt.Println("delayFun")
}
