package main

import (
	"fmt"
	"sync"
)

func main() {
	var sm sync.Map

	for i := 0; i < 10; i++ {
		sm.Store(i, i)
	}

	sm.Range(func(key, value interface{}) bool {
		ti := key.(int)
		//fmt.Println(key)
		if ti == 5 {
			sm.Delete(key)
		}

		return true
	})

	num := 0
	sm.Range(func(key, value interface{}) bool {
		num++
		return true
	})
	fmt.Println("total", num)
}
