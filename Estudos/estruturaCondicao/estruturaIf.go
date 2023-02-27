package main

import (
	f "fmt"
	"math/rand"
	"time"
)

func numRandom() int {
	s := rand.NewSource(time.Now().UnixNano())
	return rand.New(s).Intn(10)
}
func main() {
	// se a condição for satisfeita, printa "Ganhou"
	// Senão printa "não ganhou"
	if i := numRandom(); i > 5 {
		f.Println("Ganhou")
	} else {
		f.Println("Perdeu")
	}
}
