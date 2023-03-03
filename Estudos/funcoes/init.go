package main

import "fmt"

// por padrão a função init é executado antes da função main
func init() {
	fmt.Println("Inicializando")
}

func main() {
	fmt.Println("main...")
}
