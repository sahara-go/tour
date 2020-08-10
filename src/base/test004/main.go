package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		close(ch)
	}()
	ch <- 1
	fmt.Println("Done")
}
