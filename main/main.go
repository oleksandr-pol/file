package main

import (
	"fmt"

	"github.com/oleksandr-pol/file"
)

func main() {
	r, _ := file.IsExists("test.txt")
	fmt.Println(r)

	b, _ := file.GetContents("test.txt")
	s := string(b[:])
	fmt.Println(s)

	e := file.PutContents("test.txt", "MIT LICENCE")
	if e != nil {
		fmt.Println(e)
	}

	b1, _ := file.GetContents("test.txt")
	s1 := string(b1[:])
	fmt.Println(s1)
}
