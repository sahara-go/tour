// 单例设计模式
package main

import (
	"fmt"
	"sync"
)

func main() {
	// 测试
	TestSingleton(100)
}

func TestSingleton(parCount int) {
	wg := sync.WaitGroup{}
	wg.Add(parCount)

	instances := make([]*Singleton, parCount, parCount)
	//创建实例
	for i := 0; i < parCount; i++ {
		go func(index int) {
			instances[index] = GetInstance()
			wg.Done()
		}(i)
	}

	wg.Wait()
	for i := 0; i < parCount-1; i++ {
		if instances[i] != instances[i+1] {
			fmt.Println("instance is not equal")
		}
	}

	fmt.Println("end")
}
