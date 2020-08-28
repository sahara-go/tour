package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

type CustomError struct{}

func (*CustomError) Error() string {
	return ""
}

func main() {
	//打印反射对象的类型和值
	var s Student
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	fmt.Println(t, v)

	author := "Dave"
	fmt.Println("TypeOf author:", reflect.TypeOf(author))
	fmt.Println("ValueOf author:", reflect.ValueOf(author))

	typeOfError := reflect.TypeOf((*Student)(nil)).Elem()
	fmt.Println(reflect.TypeOf((*Student)(nil)).Kind())
	fmt.Println(reflect.TypeOf((*Student)(nil)))
	fmt.Println(typeOfError)
}
