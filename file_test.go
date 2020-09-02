package file_test

import (
	"fmt"

	"github.com/oleksandr-pol/file"
)

func Example_isExists() {
	t, _ := file.IsExists("/README.md")
	fmt.Println(t)
}

func Example_getContents() {
	c, _ := file.GetContents("/README.md")
	fmt.Println(c)
}

func Example_putContents() {
	d1 := []byte("hello\ngo\n")
	file.PutContents("/README.md", d1)
}
