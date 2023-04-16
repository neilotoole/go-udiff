package main

import (
	"fmt"

	"github.com/aymanbagabas/go-udiff"
)

func main() {
	a := "Hello, world!"
	b := "Hello, Go!"

	edits := udiff.Strings(a, b)
	d, err := udiff.ToUnifiedDiff("a.txt", "b.txt", a, edits)
	if err != nil {
		panic(err)
	}

	fmt.Println(d.String())
}