package main

import (
	"flag"
	"fmt"
)

func main() {
	strPtr := flag.String("strArg", "go", "a string")
	intPtr := flag.Int("intArg", 12, "a int")
	boolPtr := flag.Bool("boolArg", true, "a bool")

	var str2 string
	flag.StringVar(&str2, "strArg2", "hello", "a string")

	flag.Parse()
	fmt.Println("strPtr:", *strPtr)
	fmt.Println("intPtr:", *intPtr)
	fmt.Println("boolPtr", *boolPtr)
	fmt.Println("strPtr2:", str2)
}
