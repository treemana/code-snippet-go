package main

import "fmt"

func main() {
	needSort3 := []int{2, 8, 6, 5, 3, 9}
	qSort3(needSort3)
	for _, v := range needSort3 {
		fmt.Println(v)
	}

}

func qSort3(source []int) {
	n := len(source)
	if n <= 1 {
		return
	}

	l := 0
	r := n - 1
	key := source[0]

	for i := 1; i <= r; {
		if source[i] > key {
			source[r], source[i] = source[i], source[r]
			r--
			continue
		}
		source[l], source[i] = source[i], source[l]
		i++
		l++
	}
	qSort3(source[:l])
	qSort3(source[l+1:])
}
