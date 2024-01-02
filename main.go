package main

import "fmt"

const PREFIXOOLAPORTUGUES = "Olá, "

func Ola(nome string) string {
	if nome == "" {
		nome = "Mundo"
	}
	return PREFIXOOLAPORTUGUES + nome
}

func main() {
	fmt.Println(Ola("mundo"))
}
