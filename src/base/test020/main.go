package main

import "fmt"

func sum(a, b int) int {
	a2 := a * a
	b2 := b * b
	c := a2 + b2

	fmt.Println(c)

	return c
}

func main() {
	sum(1, 2)
}
