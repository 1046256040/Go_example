package main

import (
	"fmt"
	"log"
	"runtime"
)

func trace(a int) int {
	if a <= 1 {
		return 1
	} else {
		// fmt.Println(a)
		return trace(a-2) + trace(a-1)
	}
}

func factor(a int64) int64 {
	if a == 1 {
		return 1
	} else {
		return factor(a-1) * a
	}
}

func f() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

func fabo() func() int {
	var a int = 0
	var b int = 1
	return func() int {
		c := a
		a = b
		b = a + c
		return c
	}
}

// var where = log.Print

func main() {
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}

	log.SetFlags(log.Llongfile)

	fmt.Println("--- print 10 to 1 ---")
	fmt.Println("10", trace(10))

	where()

	fmt.Println("--- factor ---")
	fmt.Println(factor(15))

	fmt.Println("--- f ---")
	fmt.Println(f())

	where()

	f := fabo()
	fmt.Println("--- fabo ---")
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	fmt.Println("--- end ---")

	where()
}
