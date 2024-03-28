// main.go

package main

import (
	"fmt"
	"mathutil"
)

func main() {
	// Use as funções do pacote mathutil
	sum := mathutil.Add(10, 5)
	difference := mathutil.Subtract(10, 5)

	// Imprima os resultados
	fmt.Println("Soma:", sum)
	fmt.Println("Diferença:", difference)
}
