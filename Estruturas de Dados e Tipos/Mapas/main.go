package main

import "fmt"

func main() {
	// Mapa (também conhecido como dicionário)
	mapa := map[string]int{
		"um":   1,
		"dois": 2,
		"três": 3,
	}

	// Adicionando um novo par chave-valor
	mapa["quatro"] = 4

	fmt.Println(mapa)
}
