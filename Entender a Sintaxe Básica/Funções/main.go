package main

import "fmt"

// Função simples que retorna a soma de dois inteiros
func soma(a, b int) int {
	return a + b
}

func main() {
	resultado := soma(10, 5)
	fmt.Println("Resultado da soma:", resultado)
}
