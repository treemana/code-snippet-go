package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	var ss []int64

	for i := 0; i < 5; i++ {
		ss = append(ss, int64(i))
	}

	var str1 string
	var str1Time1 = time.Now().Nanosecond()
	var buffer bytes.Buffer
	for _, v := range ss {
		buffer.WriteString(strconv.FormatInt(v, 10))
		buffer.WriteString(",")
	}
	str1 = buffer.String()
	var str1Time2 = time.Now().Nanosecond()
	fmt.Println(str1)
	fmt.Println(str1Time2 - str1Time1)
	fmt.Println()
	replace()
	trim()
}

func replace() {
	var ss = "this is a word"
	fmt.Println(strings.Replace(ss, "is", "si", 3))
}

func trim() {
	fmt.Println("func trim()")
	var ss = "aaa"
	sss := strings.Trim(ss, "a")
	fmt.Println(sss)
}
