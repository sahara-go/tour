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
			_, ok := <-chd
			fmt.Println(v)
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
			_, ok := <-cha
			fmt.Println(string(v))
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
