package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path := "/sss/abc/sdf.zzz"

	fmt.Println(filepath.Ext(path))
}
