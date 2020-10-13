package main

import "testing"

func TestSoma(t *testing.T) {
    total := Soma(5, 1)
    if total != 10 {
       t.Errorf("Soma incorreta, valor: %d, esperado: %d.", total, 10)
    }
}