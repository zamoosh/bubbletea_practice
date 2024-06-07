package main

import (
	"fmt"
	"strings"
)

func main() {
	l := []string{"ali", "mohammad reza", "sara", "nob", "ashly"}

	// ls := fmt.Sprintf("%v", l)[1:]
	// ls = ls[:len(ls)-1]
	ls := strings.Join(l, "\n")

	fmt.Println(ls)
}
