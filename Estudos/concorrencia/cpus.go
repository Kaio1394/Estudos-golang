package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Número de processadores
	fmt.Println(runtime.NumCPU())
}
