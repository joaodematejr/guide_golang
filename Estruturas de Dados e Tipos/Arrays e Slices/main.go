package main

import "fmt"

func main() {
	// Array com tamanho fixo
	var array1 [3]int = [3]int{1, 2, 3}

	// Slice (uma visão flexível de um array)
	slice1 := []int{4, 5, 6}

	// Adicionando elementos a um slice
	slice1 = append(slice1, 7)

	fmt.Println(array1, slice1)
}
