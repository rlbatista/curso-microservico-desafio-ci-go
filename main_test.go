package main

import "testing"

func TestSoma(t *testing.T) {
    tables := []struct {
        op1 int
        op2 int
        tot int
    }{
        {1, 1, 2},
        {3, 5, 8},
        {7, 2, 9},
    }

    for _, table := range tables {
        total := Soma(table.op1, table.op2)
        if total != table.tot {
            t.Errorf("Soma (%d + %d) incorreta, esperado %d | obtido %d", table.op1, table.op2, table.tot, total)
        }
    }
            
} 
