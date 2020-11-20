package localcache

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

func Main() {
	key := "testCache"
	value := map[int64]string{123: "qqq"}
	localSnapshot := cache.New(time.Minute*5, time.Minute*10)
	localSnapshot.Set(key, value, time.Minute*5)
	t, ok := localSnapshot.Get(key)
	if ok {
		v, ok := t.(map[int64]string)
		if ok {
			for k, v := range v {
				fmt.Println("key:", k, "value:", v)
			}
			return
		}
		fmt.Println("not strings")
		return
	}
	fmt.Println("not get")
	return
}
