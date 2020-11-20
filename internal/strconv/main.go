package main

import (
	"fmt"
	"strconv"
)

func main() {
	var intSlice = []int64{2, 3, 1}
	var strBytes []byte
	for _, value := range intSlice {
		fmt.Println(value)
		strconv.AppendInt(strBytes, value, 10)
	}
	fmt.Println(strBytes)
	fmt.Println(string(strBytes))
}
