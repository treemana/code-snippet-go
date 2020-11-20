package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func fileBase64() {
	filePath := "/Users/treeman-zhou/Desktop/temp.jpeg"

	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	// Base64 Standard Encoding
	sEnc := base64.StdEncoding.EncodeToString(bytes)
	fmt.Println(sEnc)
}

func main() {
	fileBase64()
}
