package main

import (
	"fmt"
	"testing"
	"log"
)

func TestSoma(t *testing.T) {
	a := 1
    b := 5
    total := Soma(a, b)
    if total != 10 {
       t.Errorf("Soma incorreta, valor: %d, esperado: %d.", total, 10)
    } else {
		t.Logf("%d + %d = %d", a, b, total)
		fmt.Printf("%d + %d = %d\n", a, b, total)
		log.Println("5 + 5 = 10")
	}
}