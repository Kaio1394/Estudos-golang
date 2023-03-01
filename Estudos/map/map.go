package main

import "fmt"

// passando map como parâmetro de uma função
func add(estrutura map[int]string, cpf int, name string) {
	estrutura[cpf] = name
}

func main() {
	// criando um map
	aprovados := make(map[int]string)
	aprovados[1241415] = "Kaio"
	aprovados[5456465] = "Eliziane"
	aprovados[56164516] = "Maria"
	aprovados[651465451] = "Jorge"
	aprovados[54545454] = "Yuri"
	fmt.Println(aprovados)

	add(aprovados, 154151541, "José")
	for key, value := range aprovados {
		fmt.Println(key, value)
	}
}
