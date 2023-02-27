package main

import (
	f "fmt"
)

func change(ptr *int) {
	*ptr = 5
}

func main() {
	i := 1

	// Criando um ponteiro do tipo inteiro
	// Diferente da linguagem c/c++, golang não possui a aritmética de ponteiros
	var ptr *int = nil
	ptr = &i
	f.Println("Endereço de memoria da variável i ", &i)
	f.Println("Endereço de memoria da variável i associado ao ponteiro ", ptr)
	f.Println("Valor que o ponteiro está apontando ", *ptr)
	f.Println("Endereço de memoria do ponteiro ", &ptr)

	// mudando valor da variável através de um função  que tem como parâmetro
	// um ponteiro
	change(ptr)
	f.Println(i)
}
