package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scaner := bufio.NewScanner(os.Stdin)

	for scaner.Scan() {
		s := strings.ToUpper(scaner.Text())
		fmt.Println(s)
	}

	if err := scaner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "err:", err)
		os.Exit(1)
	}
}
