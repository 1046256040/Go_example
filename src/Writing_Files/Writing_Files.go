package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	d1 := []byte("na ku na ma ta ta")
	e := ioutil.WriteFile("dat.txt", d1, 644)
	checkErr(e)

	f, err := os.Create("dat2.txt")
	checkErr(err)
	defer f.Close()

	d2 := []byte("ta ta ma na ku na")
	n2, err := f.Write(d2)
	checkErr(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	d4 := []byte("ta ta ma ku ku")
	w := bufio.NewWriter(f)
	n4, err := w.Write(d4)
	checkErr(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()
}
