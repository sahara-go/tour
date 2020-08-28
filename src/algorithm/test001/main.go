// 冒泡排序算法
package main

import "fmt"

func main() {
	testArr := []int{5, 12, 9, 77, 44, 2, 55}
	bubbleSort(testArr)
	fmt.Println(testArr)
}

// bubbleSort 冒泡排序算法
func bubbleSort(arr []int) {
	if len(arr) == 0 {
		return
	}
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return
}
