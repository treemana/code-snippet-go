package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	h := sha256.New()
	h.Write([]byte("hello world\n"))
	hash := hex.EncodeToString(h.Sum(nil))
	fmt.Println(hash)
}
