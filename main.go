package main

import (
	"fmt"
	"sort"
)

func main() {

	a := []int{5, 1, 3, 4, 6, 9, 8, 7, 3}
	sort.Ints(a)
	fmt.Println(a)
	fmt.Println("vim-go")
}
