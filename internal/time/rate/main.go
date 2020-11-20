package main

import (
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

const (
	r   = 2 // 每秒生产 token 数量
	max = 5 // 桶的容量大小
)

var gameScene = rate.NewLimiter(r, max)

func main() {

	var i int
	var j int
	// 清空漏桶
	for {
		if !gameScene.Allow() {
			break
		}
		fmt.Println(fmt.Sprintf("第 %d 次循环 token 下标 %d", i, j))
		j++
		i++
	}

	fmt.Println("桶已清空")

	// 等待 3 秒, 将产生 6 个 token, 但 max 只能容纳 5 个, 丢弃一个
	time.Sleep(time.Second * 3)

	// 以下打印 5 个
	i = 0
	j = 0
	for {
		if !gameScene.Allow() {
			break
		}
		fmt.Println(fmt.Sprintf("第 %d 次循环 token 下标 %d", i, j))
		j++
		i++
	}

}
