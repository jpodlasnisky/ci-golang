package main

import "testing"

func TesteSoma(t *testing.T) {
	resultado := Soma(5, 5)
	esperado := 10

	if resultado != esperado {
		t.Errorf("Resultado %d, Esperado %d", resultado, esperado)
	}
}