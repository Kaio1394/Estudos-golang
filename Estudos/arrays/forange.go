package main

import "fmt"

func main() {
	// declarando um array
	// [...] significa que o compilador irá contar o tamanho do array
	numeros := [...]int{11, 2, 3, 41, 5}
	fmt.Print(numeros)

	// percorre o array. Equivalente ao for in
	for i, numero := range numeros {
		fmt.Println()
		fmt.Printf("%d) %d\n", i+1, numero)
	}
	// o '_' ignora o índice
	for _, numero := range numeros {
		fmt.Println()
		fmt.Println(numero)
	}

}
