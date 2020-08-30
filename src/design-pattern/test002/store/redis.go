// 单例设计模式
package store

import "fmt"

type RedisStore struct {
}

func (s *RedisStore) Save() {
	fmt.Println("redis store")
}
