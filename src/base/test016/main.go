package main

import (
	"bytes"
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
)

func main() {
	addInt32()
	unSignInt32()
}

// addInt32 演示并发累加
func addInt32() {
	var n int32
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()
			// n = n + 1 这种方式累加结果不一定等于1000
			atomic.AddInt32(&n, 1)
		}()
	}
	wg.Wait()
	fmt.Println(n)
}

func unSignInt32() {
	var (
		n  uint64 = 97
		m  uint64 = 1
		k  int    = 2
		n2 uint32 = 97
	)

	const (
		//a = 3
		b uint32 = 4
	)

	atomic.AddUint64(&n, -m)
	fmt.Println(n) // 97 -1
	atomic.AddUint64(&n, -uint64(k))
	fmt.Println(n) // 96 - 2
	atomic.AddUint32(&n2, ^(b - 1))
	fmt.Println(n) // 94 - 4
	b2 := ^(b - 1)
	fmt.Println(b2)
	fmt.Println(reflect.TypeOf(b2))
	fmt.Println(coverInt32ToBitStr(b2))
}

func coverInt32ToBitStr(n uint32) string {
	var rs bytes.Buffer
	var m uint32 = 1
	for i := 0; i < 32; i++ {
		if ((n >> i) & m) == 1 {
			rs.WriteString("1")
		} else {
			rs.WriteString("0")
		}
	}
	rsBytes := rs.Bytes()
	l := len(rsBytes)
	for i := 0; i <= l/2; i++ {
		rsBytes[i], rsBytes[l-i-1] = rsBytes[l-i-1], rsBytes[i]
	}

	return string(rsBytes)
}
