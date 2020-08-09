package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	chd := make(chan bool, 1)
	cha := make(chan bool, 1)
	var digestArr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var charArr = "abcdefghigk"
	wg.Add(2)
	go func() {
		defer func() {
			wg.Done()
		}()
		for _, v := range digestArr {
			// 阻塞等待，直到chd有值，或者chd已关闭，都会打印v值
			_, ok := <-chd
			fmt.Println(v)
			// 如果chd已关闭，会一直返回默认值false, 判断如果chd已经被关闭，
			// 就不再往cha 写数据
			if ok {
				cha <- true
			}
		}
		close(cha)
	}()

	go func() {
		defer func() {
			wg.Done()
		}()
		for _, v := range []byte(charArr) {
			// 阻塞等待，直到cha有值，或者cha已关闭，都会打印v值
			_, ok := <-cha
			fmt.Println(string(v))
			// 如果cha已关闭，会一直返回默认值false, 判断如果cha已经被关闭，
			// 就不再往chd 写数据
			if ok {
				chd <- true
			}
		}
		close(chd)
	}()
	chd <- true
	wg.Wait()
	fmt.Println("Done")
}
