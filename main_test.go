package main

import "fmt"
import "testing"

func TestSoma(t *testing.T) {
	a := 5
    b := 5
    total := Soma(a, b)
    if total != 10 {
       t.Errorf("Soma incorreta, valor: %d, esperado: %d.", total, 10)
    } else {
		t.Logf("%d + %d = %d", a, b, total)
		fmt.Printf("%d + %d = %d\n", a, b, total)
	}
}