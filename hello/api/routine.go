package api

import (
	"fmt"
	"time"
)

func Running() {
	var times int
	// 构建一个无限循环
	for {
		times++
		fmt.Println("tick", times)

		// 延时1s
		time.Sleep(time.Second)
	}
}
