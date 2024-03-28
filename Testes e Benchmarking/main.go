package main

import (
	"testing"
)

// Função para somar todos os números em uma fatia de inteiros.
func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// Teste unitário para a função Sum.
func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	expected := 15
	result := Sum(numbers)
	if result != expected {
		t.Errorf("Sum(%v) = %d; esperado %d", numbers, result, expected)
	}
}

// Benchmark para a função Sum.
func BenchmarkSum(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}
