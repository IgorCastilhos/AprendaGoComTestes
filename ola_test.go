package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("Igor")
	esperado := "Olá, Igor"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
