package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("---------------- slice sort start")
	var sliceSort = []int64{2, 3, 1}
	sort.Slice(sliceSort, func(i, j int) bool {
		return sliceSort[i] < sliceSort[j]
	})
	for _, s := range sliceSort {
		fmt.Println(s)
	}
	fmt.Println("---------------- slice sort end")

	var ss1 []string
	var ss2 = make([]string, 0)
	ss1 = append(ss1, ss2...)
	fmt.Println(len(ss1))
	fmt.Println(ss1)
	chop()
}

func chop() {
	ss := []string{"a", "b", "c"}
	fmt.Println(ss[0:1])
	fmt.Println(ss[1:])
}
