package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"a", "b", "c"}
	sort.Strings(strs)
	fmt.Println("strs:", strs)

	ints := []int{3, 1, 2}
	sort.Ints(ints)
	fmt.Println("ints:", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("sorted:", s)
}
