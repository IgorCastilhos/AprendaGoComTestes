package main

import "fmt"

const (
	ESPANHOL            = "espanhol"
	FRANCES             = "francês"
	PREFIXOOLAPORTUGUES = "Olá, "
	PREFIXOOLAESPANHOL  = "Hola, "
	PREFIXOOLAFRANCES   = "Bonjour, "
)

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}
	return prefixodeSaudacao(idioma) + nome
}

func prefixodeSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case FRANCES:
		prefixo = PREFIXOOLAFRANCES
	case ESPANHOL:
		prefixo = PREFIXOOLAESPANHOL
	default:
		prefixo = PREFIXOOLAPORTUGUES
	}
	return
}

func main() {
	fmt.Println(Ola("mundo", ""))
}
