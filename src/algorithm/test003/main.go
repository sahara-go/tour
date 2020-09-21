// 快速排序法
package main

import "fmt"

func main() {
	testArr := []int{5, 12, 9, 3, 77, 44, 2, 55, 4, 66}
	QuickSort(testArr)
	fmt.Println(testArr)
}

// QuickSort 快速排序
func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	startIndex := 0
	left := startIndex
	right := len(arr) - 1
	pivot := arr[startIndex]

	for left < right {
		for left < right && pivot < arr[right] {
			right--
		}

		for left < right && pivot >= arr[left] {
			left++
		}

		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}

	}

	arr[startIndex], arr[left] = arr[left], arr[startIndex]
	QuickSort(arr[0:left])
	QuickSort(arr[left+1:])
	return
}

type Interface interface {
	Fun()
}

type A struct {
}

func (t *A) Fun(v Interface) {
	v.Fun()
}
