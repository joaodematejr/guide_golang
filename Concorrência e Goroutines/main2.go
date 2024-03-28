package main

import (
	"fmt"
	"time"
)

func tarefa(id int) {
	// Simula uma tarefa demorada
	fmt.Printf("Tarefa %d iniciada\n", id)
	time.Sleep(time.Second * 2)
	fmt.Printf("Tarefa %d concluída\n", id)
}

func main() {
	// Executando tarefas sequencialmente
	for i := 1; i <= 3; i++ {
		tarefa(i)
	}

	// Executando tarefas concorrentemente usando goroutines
	for i := 1; i <= 3; i++ {
		go tarefa(i)
	}

	// Aguarda um pouco para permitir que as goroutines sejam executadas
	time.Sleep(time.Second * 3)
}
