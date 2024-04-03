package main

import (
	"math/rand"
	"testing"
)

// Função para calcular a soma de uma lista de números inteiros
func somaLista(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// Benchmark para a função somaLista
func BenchmarkSomaLista(b *testing.B) {
	// Gere uma lista de números aleatórios para testar
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = rand.Intn(100)
	}

	// Execute a função somaLista b.N vezes e meça o tempo de execução
	for i := 0; i < b.N; i++ {
		somaLista(nums)
	}
}
