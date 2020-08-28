package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	go recoverTest1()
	time.Sleep(time.Second * 3)
	fmt.Println("正常结束")
}

// recoverTest1 协程panic处理
func recoverTest1() {
	defer func() {
		//协程内包含协程的时候，每个协程都要处理panic，否则会导致进程退出
		if err := recover(); err != nil {
			fmt.Printf("panic recover : %v", err)
		}
	}()
	go recoverTest2()
}

func recoverTest2() {
	defer func() {
		//没有这段代码，父协程无法捕获panic,程序会直接退出
		if err := recover(); err != nil {
			fmt.Printf("panic recover2 : %v", err)
		}
	}()
	time.Sleep(time.Second * 1)
	panic(errors.New("recoverTest2 panic"))
}
