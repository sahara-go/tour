package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

type Foo struct {
	name string
	age  int
}

func main() {
	//new(T) 为一个 T 类型新值分配空间并将此空间初始化为 T 的零值，返回的是新值的地址，也就是 T 类型的指针 *T，该指针指向 T 的新分配的零值
	fmt.Println("new例子")
	stu := new(Student)
	fmt.Println(stu)
	stu2 := &Student{}
	fmt.Println(stu2)
	//以下代码演示了 struct 初始化的过程，可以说明不使用 new 一样可以完成 struct 的初始化工作
	fmt.Println("以下代码演示了 struct 初始化的过程，可以说明不使用 new 一样可以完成 struct 的初始化工作")
	//声明初始化
	var foo1 Foo
	fmt.Printf("foo1 --> %#v\n ", foo1) //main.Foo{age:0, name:""}
	foo1.age = 1
	fmt.Println(foo1.age)

	//struct literal 初始化
	foo2 := Foo{}
	fmt.Printf("foo2 --> %#v\n ", foo2) //main.Foo{age:0, name:""}
	foo2.age = 2
	fmt.Println(foo2.age)

	//指针初始化
	foo3 := &Foo{}
	fmt.Printf("foo3 --> %#v\n ", foo3) //&main.Foo{age:0, name:""}
	foo3.age = 3
	fmt.Println(foo3.age)

	//new 初始化
	foo4 := new(Foo)
	fmt.Printf("foo4 --> %#v\n ", foo4) //&main.Foo{age:0, name:""}
	foo4.age = 4
	fmt.Println(foo4.age)

	//声明指针并用 new 初始化
	var foo5 *Foo = new(Foo)
	fmt.Printf("foo5 --> %#v\n ", foo5) //&main.Foo{age:0, name:""}
	foo5.age = 5
	fmt.Println(foo5.age)

	fmt.Println("make例子")
	var m1 map[int]string
	// m1 is always not nil
	if m1 == nil {
		fmt.Printf("m1 is nil --> %#v \n ", m1) //map[int]string(nil)
	}

	m2 := make(map[int]string)
	// m2 is always not nil
	if m2 == nil {
		fmt.Printf("m2 is nil --> %#v \n ", m2)
	} else {
		fmt.Printf("m2 is not nill --> %#v \n ", m2)
	}

	var c1 chan string
	// c1 is always not nil
	if c1 == nil {
		fmt.Printf("c1 is nil --> %#v \n ", c1) //(chan string)(nil)
	}

	c2 := make(chan string)
	// c2 is always not nil
	if c2 == nil {
		fmt.Printf("c2 is nil --> %#v \n ", c2)
	} else {
		fmt.Printf("c2 is not nill --> %#v \n ", c2) //(chan string)(0xc420016120)
	}
}
