// 简单工厂设计模式
package main

import "github.com/sahara-gopher/tour/src/design-pattern/test002/store"

// Factory 工厂类
type Factory struct {
}

func (f *Factory) Create(name string) store.Store {
	switch name {
	case "redis":
		return &store.RedisStore{}
	case "mysql":
		return &store.MysqlStore{}
	default:
		return nil
	}
}

// 测试简单工厂模式
func main() {
	//创建工厂实例
	f := new(Factory)
	s := f.Create("redis")
	s.Save()
	s = f.Create("mysql")
	s.Save()
}
