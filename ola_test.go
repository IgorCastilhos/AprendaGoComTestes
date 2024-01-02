package main

import "testing"

func TestOla(t *testing.T) {
	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resultado := Ola("Igor")
		esperado := "Olá, Igor"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz 'Olá, mundo' quando uma string vazia for passada", func(t *testing.T) {
		resultado := Ola("")
		esperado := "Olá, Mundo"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	resultado := Ola("Igor")
	esperado := "Olá, Igor"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
