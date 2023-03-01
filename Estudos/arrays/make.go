package main

import "fmt"

func main() {
	s := make([]int, 10)
	s[9] = 10
	fmt.Println(s)

	// make(type, len, capacity)
	// criação de um slice com tamanho 10, mas com capacidade igual 20
	s = make([]int, 10, 20)

	// camp() -> retorna a capacidade do slice
	fmt.Println(cap(s))

	// Adicionando mais elemento no slice
	s = append(s, 1, 2, 4, 5, 6, 7, 8, 9, 10, 12, 13)
	fmt.Println(s)

}
