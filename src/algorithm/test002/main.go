// 选择序算法
package main

import "fmt"

func main() {
	testArr := []int{5, 12, 9, 77, 44, 2, 55, 66}
	bubbleSort(testArr)
	fmt.Println(testArr)
}

// bubbleSort 冒泡排序算法
func bubbleSort(arr []int) {
	if len(arr) == 0 {
		return
	}
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j <= len(arr)-1; j++ {
			if arr[j] < arr[minIndex] {
				//记录本次循环最小值位置
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return
}
