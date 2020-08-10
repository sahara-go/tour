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
}
