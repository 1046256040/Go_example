/*
   fibonacci在内存中缓存每次fibonacci的值，来提升计算速度，总耗时只需1ms
   fibonacci2是最常见的方法，每次都从头计算，效率很低，总耗时1.592s
*/
package main

import (
	"fmt"
	"time"
)

const (
	LIM = 41
)

var fibs [LIM]int64

func fibonacci(i int) (res int64) {
	if fibs[i] != 0 {
		res = fibs[i]
		return
	}
	if i <= 1 {
		res = 1
	} else {
		res = fibonacci(i-1) + fibonacci(i-2)
	}
	fibs[i] = res
	return
}

func fibonacci2(i int) (res int64) {
	if i <= 1 {
		return 1
	} else {
		return fibonacci2(i-1) + fibonacci2(i-2)
	}
}

func main() {
	var result int64 = 0
	begin := time.Now()
	for i := 1; i < LIM; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	sub := end.Sub(begin)
	fmt.Printf("fibonacci cost time:%s\n", sub)

	begin = time.Now()
	for i := 1; i < LIM; i++ {
		result = fibonacci2(i)
		fmt.Printf("fibonacci2(%d) is: %d\n", i, result)
	}
	end = time.Now()
	sub = end.Sub(begin)
	fmt.Printf("fibonacci2 cost time:%s\n", sub)
}
