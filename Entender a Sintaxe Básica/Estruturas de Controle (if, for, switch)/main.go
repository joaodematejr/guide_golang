package main

import "fmt"

func main() {
	// Estrutura if
	idade := 20
	if idade >= 18 {
		fmt.Println("Você é maior de idade.")
	} else {
		fmt.Println("Você é menor de idade.")
	}

	// Estrutura for
	for i := 0; i < 5; i++ {
		fmt.Println("Número:", i)
	}

	// Estrutura switch
	dia := "quarta-feira"
	switch dia {
	case "segunda-feira":
		fmt.Println("Hoje é segunda-feira.")
	case "terça-feira":
		fmt.Println("Hoje é terça-feira.")
	case "quarta-feira":
		fmt.Println("Hoje é quarta-feira.")
	default:
		fmt.Println("Não é segunda, terça ou quarta-feira.")
	}
}
