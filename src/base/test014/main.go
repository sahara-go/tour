package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	str := "hello world!"
	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	b := *(*[]byte)(unsafe.Pointer(strHeader))
	fmt.Println(b)

}
