package main

import "fmt"

func main() {
	// declarando um array
	// [...] significa que o compilador irá contar o tamanho do array
	numeros := [...]int{1, 2, 3, 4, 5}
	fmt.Print(numeros)
}
