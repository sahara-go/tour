// 单例设计模式
package main

import "sync"

var (
	instance *Singleton
	once     sync.Once
)

type Singleton struct {
}

// GetInstance 获取实例
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}
