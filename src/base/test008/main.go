package main

import (
	"fmt"
	"strings"
)

type StoreType string

func main() {
	var b StoreType
	b = "aa"
	strFunc(b)
}

func strFunc(str string) {
	rs := strings.Compare("a", str)
	fmt.Println(rs)
}