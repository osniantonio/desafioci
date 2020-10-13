package main

import "fmt"

func soma(a int, b int) int {
    return a + b
}

func main() {
	a := 5
    b := 5
	total := soma(a, a)
	fmt.Printf("%d + %d = %d\n", a, b, total)
}