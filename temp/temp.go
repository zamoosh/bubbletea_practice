package main

import (
	"fmt"
	"strings"
)

func main() {
	l := []string{"ali", "mohammad reza", "sara", "nob", "ashly"}

	cl := make([]string, 0, 10)
	for i, item := range l {
		cl = append(cl, fmt.Sprintf("[%d] %s", i, item))
	}
	ls := strings.Join(cl, "\n")

	fmt.Println(ls)
}
