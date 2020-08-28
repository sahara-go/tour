package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	// 创建通道并初始化
	fmt.Printf("%#v", ch)
	go func() {
		time.Sleep(3 * time.Second)
		close(ch)
	}()
	ch <- 1
	fmt.Println("Done")
}
