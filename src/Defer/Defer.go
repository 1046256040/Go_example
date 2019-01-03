package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("file.txt")
	writeFile(f)
	defer closeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("----- createFile -----")

	f, err := os.Create("file.txt")
	if err != nil {
		panic("createFile error")
	}

	return f
}

func writeFile(f *os.File) {
	fmt.Println("----- writeFile -----")

	if f != nil {
		fmt.Fprintln(f, "data")
	}
}

func closeFile(f *os.File) {
	fmt.Println("----- closeFile -----")

	if f != nil {
		f.Close()
		f = nil
	}
}
