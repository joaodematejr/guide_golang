package main

import (
	"fmt"
)

// função para calcular a soma dos quadrados de um conjunto de números
func sumOfSquares(numbers []int, resultChan chan int) {
	sum := 0
	for _, num := range numbers {
		sum += num * num
	}
	// envia o resultado para o canal
	resultChan <- sum
}

func main() {
	// conjunto de números para calcular a soma dos quadrados
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// cria um canal para receber os resultados
	resultChan := make(chan int)

	// divide os números em duas partes para processamento em paralelo
	halfLength := len(numbers) / 2
	firstHalf := numbers[:halfLength]
	secondHalf := numbers[halfLength:]

	// inicia duas goroutines para calcular a soma dos quadrados de cada metade dos números
	go sumOfSquares(firstHalf, resultChan)
	go sumOfSquares(secondHalf, resultChan)

	// aguarda a conclusão de ambas as goroutines e recebe os resultados
	result1 := <-resultChan
	result2 := <-resultChan

	// calcula a soma total dos quadrados
	totalSum := result1 + result2

	fmt.Println("A soma dos quadrados é:", totalSum)
}
