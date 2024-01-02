package main

import "fmt"

func Ola(nome string) string {
	return fmt.Sprintf("Olá, %s", nome)
}

func main() {
	fmt.Println(Ola("mundo"))
}
