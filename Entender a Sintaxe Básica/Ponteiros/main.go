package main

import "fmt"

func main() {
	x := 10
	fmt.Println("Valor de x:", x)
	fmt.Println("Endereço de x:", &x)

	// Declarando um ponteiro para x
	var ponteiro *int
	ponteiro = &x

	// Alterando o valor de x indiretamente através do ponteiro
	*ponteiro = 20
	fmt.Println("Novo valor de x:", x)
}
