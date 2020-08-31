// 单例设计模式
package store

import "fmt"

type MysqlStore struct {
	BaseStore
}

func (s *MysqlStore) Save() {
	fmt.Println("mysql store")
}
