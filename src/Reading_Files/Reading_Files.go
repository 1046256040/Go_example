package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func closeFile(f *os.File) {
	if f != nil {
		f.Close()
		fmt.Printf("close file")
	}
}

func main() {
	dat, err := ioutil.ReadFile("dat.txt")
	checkErr(err)
	fmt.Println(string(dat))

	f, err := os.Open("dat.txt")
	checkErr(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	checkErr(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	o2, err := f.Seek(6, 0)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	checkErr(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	o3, err := f.Seek(6, 0)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	checkErr(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	checkErr(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(4)
	checkErr(err)
	fmt.Printf("4 bytes: %s\n", string(b4))

	defer closeFile(f)

}
