package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Object struct {
	Num int
	S   []string
}

type Student struct {
	Name string
}

func main() {
	//不内存复制的情况下将string转成[]byte
	str := "hello world"
	strH := (*reflect.StringHeader)(unsafe.Pointer(&str))
	sliceH := reflect.SliceHeader{
		Data: strH.Data,
		Len:  strH.Len,
		Cap:  strH.Len,
	}
	rs := (*[]byte)(unsafe.Pointer(&sliceH))
	fmt.Println(rs)

	pa := Object{
		Num: 20,
		S: []string{
			"a",
		},
	}
	fmt.Println(unsafe.Sizeof(pa))
	fmt.Println(unsafe.Sizeof(pa.Num))
	fmt.Println(unsafe.Sizeof(pa.S))
	fmt.Println(unsafe.Offsetof(pa.Num))
	fmt.Println(unsafe.Offsetof(pa.S))
}
