package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("boilingF is %g, freezingF is %g\n", boilingF, freezingF)
	fmt.Println("vim-go")
}

func ftoc(f float64) float64 {
	return (f - 32) * 5 / 9
}
