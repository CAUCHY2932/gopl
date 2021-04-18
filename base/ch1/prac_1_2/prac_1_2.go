package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	parseWithTime()
	fmt.Println("vim-go")
}

func parseWithTime() {
	start := time.Now()
	for i, v := range os.Args[1:] {
		fmt.Println("i:", i, ",v:", v)

	}
	fmt.Printf("%d elapsed\n", time.Since(start).Microseconds())
}

func parse() {
	for i, v := range os.Args[1:] {
		fmt.Println("i:", i, ",v:", v)

	}
}
