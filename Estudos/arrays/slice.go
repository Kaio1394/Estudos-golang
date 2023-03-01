package main

import (
	"fmt"
	"reflect"
)

func main() {
	// array
	a1 := [...]int{1, 2, 3}

	// slice -> trecho de um array
	// slice não é um array, é um estrutura que aponta para o primeiro elemento
	// do array
	s1 := []int{1, 2, 3}

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[1:3]
	fmt.Println("Endereço de memória do primeiro elemento do array", &a2[0])
	fmt.Println(&s2[0])
	fmt.Println(len(s2))

}
