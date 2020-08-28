package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {
	m := make(map[string]*Student)
	stus := [3]Student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	for _, stu := range stus {
		fmt.Printf("%p", &stu)
		m[stu.Name] = &stu
	}
	fmt.Printf("%v", m["zhou"])
	fmt.Printf("%v", m["li"])

	stu := Student{
		Name: "张磊",
	}
	echoName(stu)
	var stuPtr *Student
	echoName(stuPtr)
}

func echoName(v interface{}) {
	switch t := v.(type) {
	case nil:
		fmt.Println("empty")
	case Student:
		fmt.Println(t)
	default:
		fmt.Println("no match!")
	}
}

type People struct {
	Name string
}

func (p *People) String() string {
	//fmt.Printf("%v", p)
	return fmt.Sprintf("print:%v", p)
}
