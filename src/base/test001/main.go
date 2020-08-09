package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	chd := make(chan bool, 100)
	cha := make(chan bool, 100)
	var digestArr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var charArr = "abcdefghigk"
	wg.Add(2)
	go func() {
		defer func() {
			wg.Done()
		}()
		for _, v := range digestArr {
			<-chd
			fmt.Println(v)
			cha <- true
		}
		close(cha)
	}()

	go func() {
		defer func() {
			wg.Done()
		}()
		for _, v := range []byte(charArr) {
			<-cha
			fmt.Println(string(v))
			chd <- true
		}
		close(chd)
	}()
	chd <- true
	wg.Wait()
	fmt.Println("Done")
}
