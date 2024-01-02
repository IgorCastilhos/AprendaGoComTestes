package test

import (
	"github.com/IgorCastilhos/Ola"
	"testing"
)

func TestOla(t *testing.T) {
	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resultado := main.Ola("Igor", "")
		esperado := "Olá, Igor"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz 'Olá, mundo' quando uma string vazia for passada", func(t *testing.T) {
		resultado := main.Ola("", "")
		esperado := "Olá, Mundo"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em espanhol", func(t *testing.T) {
		resultado := main.Ola("Elodie", "espanhol")
		esperado := "Hola, Elodie"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em francês", func(t *testing.T) {
		resultado := main.Ola("Mundo", "francês")
		esperado := "Bonjour, Mundo"
		verificaMensagemCorreta(t, resultado, esperado)
	})
}
