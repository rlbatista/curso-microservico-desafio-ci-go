package main

import "testing"

func TestSoma(t *testing.T) {
    soma := Soma(3, 5)
    if soma != 8 {
        t.Errorf("Soma (3 + 5) incorreta, esperado 8 | obtido %d", soma)
    }
} 
