// 工厂方法设计模式
package main

import "github.com/sahara-gopher/tour/src/design-pattern/test003/store"

// StoreFactory 工厂接口
type StoreFactory interface {
	Create() store.Store
}

// RedisStoreFactory redis store工厂
type RedisStoreFactory struct {
}

func (f *RedisStoreFactory) Create() store.Store {
	s := &store.RedisStore{}
	return s
}

// MysqlStoreFactory mysql store工厂
type MysqlStoreFactory struct {
}

func (f *MysqlStoreFactory) Create() store.Store {
	s := &store.MysqlStore{}
	return s
}

// 测试工厂方法设计模式
func main() {
	rsFactory := RedisStoreFactory{}
	msFactory := MysqlStoreFactory{}
	r := rsFactory.Create()
	r.SetHost("192.168.1.1:6379")
	r.Save()

	r = msFactory.Create()
	r.SetHost("192.168.1.2:3366")
	r.Save()
}
