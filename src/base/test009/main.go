package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	parentCtx := context.Background()
	fmt.Println(parentCtx.Done())

	ctx, cancel := context.WithTimeout(parentCtx, 1*time.Second)
	ctx2, cancel2 := context.WithTimeout(ctx, 1*time.Second)
	fmt.Println(ctx.Done())
	fmt.Println(ctx2.Done())
	defer func() {
		cancel()
		cancel2()
	}()

	//go handle(ctx, 1000*time.Millisecond)
	go cancelCtx(cancel, 10*time.Millisecond)
	select {
	case <-ctx2.Done():
		fmt.Println("handle ctx cancel:", ctx2.Err())
	}

	fmt.Printf("time:%d", time.Now().Sub(startTime).Milliseconds())
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle :", ctx.Err())
	}
}

func cancelCtx(cancel context.CancelFunc, duration time.Duration) {
	select {
	case <-time.After(duration):
		cancel()
	}
}
