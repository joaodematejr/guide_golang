package main

import "fmt"

// Definindo uma struct
type Pessoa struct {
	Nome     string
	Idade    int
	Endereco Endereco
}

// Struct aninhada
type Endereco struct {
	Rua    string
	Cidade string
}

func main() {
	// Criando uma instância de Pessoa
	pessoa1 := Pessoa{
		Nome:  "Alice",
		Idade: 30,
		Endereco: Endereco{
			Rua:    "Rua das Flores",
			Cidade: "Cidade A",
		},
	}

	fmt.Println(pessoa1)
}
