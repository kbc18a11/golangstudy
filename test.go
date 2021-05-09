package main

import (
	"fmt"
)

func main() {
	i := []int{1, 2, 3}
	p := &i
	fmt.Println(p[0])
	fmt.Println(i)
}
