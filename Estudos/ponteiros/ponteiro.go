package main

import (
	f "fmt"
)

func main() {
	i := 1

	// Criando um ponteiro do tipo inteiro
	var ptr *int = nil
	ptr = &i
	f.Println("Endereço de memoria da variável i ", &i)
	f.Println("Valor que o ponteiro está apontando ", *ptr)
	f.Println("Endereço de memoria do ponteiro ", &ptr)
}
