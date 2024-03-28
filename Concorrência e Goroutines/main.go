package main

import (
	"fmt"
	"time"
)

// Função que simula um processamento demorado
func processamentoDemorado(id int) {
	fmt.Printf("Goroutine %d: Iniciando processamento demorado...\n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d: Processamento demorado concluído.\n", id)
}

func main() {
	// Iniciar três goroutines concorrentes
	for i := 1; i <= 3; i++ {
		go processamentoDemorado(i)
	}

	// Esperar um pouco para garantir que as goroutines tenham tempo de serem executadas
	time.Sleep(3 * time.Second)

	fmt.Println("Programa principal: Todas as goroutines terminaram.")
}
