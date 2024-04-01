package main

import (
	"fmt"
)

// Definindo um tipo de função para operações sobre inteiros
type IntOperation func(int) int

// Função de ordem superior que aplica uma operação sobre cada elemento de um slice de inteiros
func MapInts(values []int, op IntOperation) []int {
	result := make([]int, len(values))
	for i, v := range values {
		result[i] = op(v)
	}
	return result
}

// Função de ordem superior que filtra elementos de um slice de inteiros com base em uma condição
func FilterInts(values []int, condition func(int) bool) []int {
	result := []int{}
	for _, v := range values {
		if condition(v) {
			result = append(result, v)
		}
	}
	return result
}

// Função de ordem superior que reduz um slice de inteiros a um único valor
func ReduceInts(values []int, initial int, reducer func(int, int) int) int {
	result := initial
	for _, v := range values {
		result = reducer(result, v)
	}
	return result
}

func main() {
	// Um slice de números inteiros
	numbers := []int{1, 2, 3, 4, 5}

	// Mapeando os números para seus quadrados
	squared := MapInts(numbers, func(x int) int {
		return x * x
	})
	fmt.Println("Squared numbers:", squared)

	// Filtrando os números pares
	even := FilterInts(numbers, func(x int) bool {
		return x%2 == 0
	})
	fmt.Println("Even numbers:", even)

	// Reduzindo os números para sua soma
	sum := ReduceInts(numbers, 0, func(acc, x int) int {
		return acc + x
	})
	fmt.Println("Sum of numbers:", sum)
}
