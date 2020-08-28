package main

import (
	"fmt"
)

type StoreType string

const (
	RedisStoreType StoreType = "redis"
	MemStoreType   StoreType = "memory"
	MysqlStoreType StoreType = "mysql"
)

func main() {
	rs := IsInStoreTypeArr(StoreType("redis"), RedisStoreType, MemStoreType, MysqlStoreType)
	fmt.Println(rs)
}

func IsInStoreTypeArr(s StoreType, inArr ...StoreType) bool {
	for _, v := range inArr {
		if v == s {
			return true
		}
	}
	return false
}
