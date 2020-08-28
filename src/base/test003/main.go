package main

import "fmt"

func main() {
	// copy操作
	s := []int{0, 1, 2, 3, 4}
	copy(s[1:], s[2:])
	fmt.Println(s)
	fmt.Println()
	// result
	// [0 2 3 4 4]

	// 截取操作
	s = []int{0, 1, 2, 3, 4}
	s = s[1:]
	fmt.Println(s)
	s = s[1:2]
	fmt.Println(s)
	fmt.Println()

	// append 函数一不小心就会踩的坑
	var str = make([]string, 10)
	str = append(str, "hello")
	fmt.Printf("s[0]=%s\n, len:%d", str[0], len(str))
	//result
	//s=[          hello],s[0]=

	fmt.Printf("str切片内存地址：%p, len:%d, cap:%d\n", str, len(str), cap(str))
	str2 := append(str, "w", "o", "w", "o", "w", "o", "w", "o", "w", "o")
	fmt.Printf("str2切片内存地址：%p, len:%d, cap:%d\n", str2, len(str2), cap(str2))
}
