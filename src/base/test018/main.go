package main

import (
	"context"
	//"encoding/json"
	"fmt"
	"golang.org/x/time/rate"
	//"reflect"
	"time"
	//"unsafe"
)

func main() {
	startTime := time.Now()
	//每1ms生成一个令牌
	limit := rate.Every(1 * time.Millisecond)
	limiter := rate.NewLimiter(limit, 100)
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer func() {
		cancel()
	}()
	for i := 0; i < 1000; i++ {
		//阻塞等待
		err := limiter.Wait(ctx)
		if err != nil {
			fmt.Printf("error:%s", err)
			break
		}
		fmt.Println("get token")
	}
	d := time.Since(startTime)
	fmt.Printf("结束,耗时：%d ms", d.Milliseconds())
}
