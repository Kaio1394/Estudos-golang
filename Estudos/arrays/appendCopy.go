package main

import "fmt"

func main() {
	array := [3]int{1, 2, 3}
	var slice []int
	// slice = [4, 5, 6]
	slice = append(slice, 4, 5, 6)
	fmt.Printf("array: %d. slice: %d", array, slice)

	// copiando os elementos do slice para o slice2
	// como o slice2 está com o tamanho limitado a 2 elementos
	// então só sera copiado os dosi primeiros elementos
	slice2 := make([]int, 2)
	copy(slice2, slice)
	fmt.Printf("\nslice: %d. slice2: %d", slice, slice2)

}
