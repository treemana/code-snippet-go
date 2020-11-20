package main

import (
	real "crypto/rand"
	"fmt"
	"math/big"
	fake "math/rand"
	"time"
)

func main() {
	fmt.Println("fake random : ", FakeRandom(100))
	fmt.Println("real random : ", RealRandom(100))
}

func FakeRandom(max int64) (r int64) {
	fake.Seed(time.Now().UnixNano())
	r = fake.Int63n(max)
	return
}

func RealRandom(max int64) (r int64) {
	result, err := real.Int(real.Reader, big.NewInt(max))
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}
	r = result.Int64()
	return
}
