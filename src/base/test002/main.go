package main

import "fmt"

func main() {
	// append会改变slice a的底层结构
	a := []int{1, 2, 3}
	for k, v := range a {
		if k == 0 {
			a = append(a, 4)
			a[1] = 3
		}
		fmt.Println(k, v)
	}
	fmt.Println(a)
	fmt.Println()
	// result:
	// 1 1
	// 2 2
	// 3 3
	// [1 3 3 4]

	fmt.Println("\nslice 删除元素，迭代结果不受影响")
	a2 := []int{0, 1, 2, 3, 4}
	for i, v := range a2 {
		fmt.Println(i, v)
		if i == 0 {
			a2 = a2[1:]
		}
	}
	fmt.Println()

	// result：
	// 0 0
	// 1 2
	// 2 3
	// 3 4
	// 4 4

	// 在slice中增加元素，会更改slice含有的元素，但不会更改遍历次数
	a3 := []int{0, 1, 2, 3, 4}
	for i, v := range a3 {
		fmt.Println(i, v)
		if i == 0 {
			a3 = append(a3, 6)
		}
	}
	fmt.Println()
	// result:
	// 0 0
	// 1 1
	// 2 2
	// 3 3
	// 4 4

	// 删除slice中的元素
	a2 = []int{0, 1, 2, 3, 4}
	for i, v := range a2 {
		fmt.Println(i, v)
		if i == 0 {
			a2 = a2[1:]
		}
	}
	fmt.Println()
	// result
	// 0 0
	// 1 1
	// 2 2
	// 3 3
	// 4 4

	// 并更改slice长度
	a2 = []int{0, 1, 2, 3, 4}
	for i, v := range a2 {
		fmt.Println(i, v)
		if i == 0 {
			a2 = a2[:len(a2)-2] //将a2的len设置为3，但并不会影响临时slice-range_temp
		}
	}
	// result
	// 0 0
	// 1 2
	// 2 3
	// 3 4
	// 4 4

	// 在遍历中删除元素
	// map迭代的时候每次输出的顺序是随机的
	fmt.Println("\n在遍历中删除元素")
	m := map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5}
	del := false
	for k, v := range m {
		fmt.Println(k, v)
		if !del {
			// 如果key 2的值在前面没有迭代到的情况下，删除后将不能再取到值
			delete(m, 2)
			del = true
		}
	}
	// result
	// 4 4
	// 5 5
	// 1 1
	// 3 3
}
