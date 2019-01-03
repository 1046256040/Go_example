package main

import (
	"fmt"
	"sort"
)

type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

/*
type byStr string

func (s byStr) Len() int {
	return len(s)
}

func (s byStr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byStr) Less(i, j int) bool {
	return s[i] < s[j]
}
*/

func main() {
	strs := []string{"kwaii", "sodesline", "yamadier"}
	sort.Sort(byLength(strs))

	fmt.Println("After sort:", strs)
}
