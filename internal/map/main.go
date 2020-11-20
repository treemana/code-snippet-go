package main

import (
	"fmt"
	"time"
)

func main() {
	var mm = map[string]string{"a": "b", "c": "d"}
	go func() {
		time.Sleep(time.Second)
		mm["e"] = "f"
	}()
	for e := range mm {
		if e == "a" {
			delete(mm, "a")
			time.Sleep(time.Second * 2)
		}
		fmt.Println(e)
	}
	fmt.Println("--------")
	for e := range mm {
		fmt.Println(e)
	}

	fmt.Println("--------make")
	var mm1 = make(map[string]string)
	var mm2 = make(map[string]string, 20)
	fmt.Println(len(mm1))
	fmt.Println(len(mm2))
	NilOrLength()
}

func NilOrLength() {
	var nilMap map[int64]string
	fmt.Println("----------NilOrLength")
	fmt.Println(nilMap == nil)
	fmt.Println(len(nilMap))
	fmt.Println("----------NilOrLength")
}
