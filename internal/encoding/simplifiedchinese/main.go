package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"unicode/utf8"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func main() {
	s1 := "中文"
	s2 := "1ab2"

	h0 := md5.New()
	h0.Write([]byte(s1 + s2))
	m0 := hex.EncodeToString(h0.Sum(nil))

	t1, _ := simplifiedchinese.GBK.NewEncoder().String(s1 + s2)
	h1 := md5.New()
	h1.Write([]byte(t1))
	m1 := hex.EncodeToString(h1.Sum(nil))

	t2, _ := simplifiedchinese.GBK.NewEncoder().String(s1)
	h2 := md5.New()
	h2.Write([]byte(t2 + s2))
	m2 := hex.EncodeToString(h2.Sum(nil))

	fmt.Println(m0)
	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m0 == m1)
	fmt.Println(m2 == m1)

	fmt.Println("t2 is utf-8 : ", utf8.ValidString(t2), t2)
	fmt.Println("s2 is utf-8 : ", utf8.ValidString(s2))
	fmt.Println("s1 is utf-8 : ", utf8.ValidString(s1))

	reader := transform.NewReader(bytes.NewReader([]byte(t2)), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		fmt.Println(e.Error())
	}
	fmt.Println(string(d))
}
