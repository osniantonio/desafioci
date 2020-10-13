package main

import "fmt"

func Soma(a int, b int) int {
    return a + b
}

func main() {
	a := 5
    b := 5
	total := Soma(a, a)
	fmt.Printf("%d + %d = %d\n", a, b, total)
}