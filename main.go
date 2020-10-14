package main

import (
	"fmt"
	"log"
)

func Soma(a int, b int) int {
	total := a + b
	fmt.Printf("%d + %d = %d\n", a, b, total)
    return total
}

func main() {
	a := 5
    b := 5
	total := Soma(a, b)
	log.Println("5 + 5 = 10")
}