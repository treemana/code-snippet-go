package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// md5File return md5sum of real file on path
func md5File(path string) (string, error) {
	var regFile *os.File
	var err error

	if regFile, err = os.Open(path); err != nil {
		return "", err
	}

	hash := md5.New()
	if _, err = io.Copy(hash, regFile); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
func main() {
	md5Str, _ := md5File("/Users/treeman-zhou/Desktop/treeman-zhou.jpg")
	fmt.Println(md5Str)
}
