package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	var num int64
	result, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}

	num = result.Int64()

	switch {
	case num > 50:
		fmt.Println("num>50:", num)
	case num < 10:
		fmt.Println("num<10:", num)
	default:
		fmt.Println(num)
	}
}
