package main

import (
	"fmt"
	"strings"
)

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vs []string, t string) bool {
	index := Index(vs, t)
	if index == -1 {
		return false
	}
	return true
}

func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	ret := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			ret = append(ret, v)
		}
	}
	return ret
}

func Map(vs []string, f func(string) string) map[int]string {
	ret := make(map[int]string)
	for i, v := range vs {
		ret[i] = f(v)
	}
	return ret
}

func main() {
	strs := []string{"peach", "peanut", "apple", "Pineapple"}

	fmt.Println(Index(strs, "apple"))
	fmt.Println(Include(strs, "banana"))

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(Filter(strs, func(v string) bool {
		return strings.HasPrefix(v, "p") || strings.HasPrefix(v, "P")
	}))

	fmt.Println(Map(strs, func(v string) string {
		return strings.ToUpper(v)
	}))
}
