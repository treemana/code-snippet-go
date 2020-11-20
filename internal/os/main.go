package main

import (
	"fmt"
	"os"
)

func main() {
	testRemove()
}

func testRemove() {
	path := "/sssssssssssss"

	err := os.Remove(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println(path + " is nil")
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println(info)
}
