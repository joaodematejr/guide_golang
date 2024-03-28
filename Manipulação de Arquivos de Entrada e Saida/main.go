package main

import (
	"fmt"
	"os"
)

func main() {
	// Criando um arquivo e escrevendo nele
	file, err := os.Create("exemplo.txt")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}
	defer file.Close()

	texto := "Olá, mundo!"
	_, err = file.WriteString(texto)
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}
	fmt.Println("Texto escrito com sucesso no arquivo.")

	// Lendo o conteúdo do arquivo
	file, err = os.Open("exemplo.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	// Criando um buffer para armazenar os dados lidos
	buffer := make([]byte, 100)
	bytesLidos, err := file.Read(buffer)
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}
	fmt.Println("Conteúdo do arquivo:")
	fmt.Println(string(buffer[:bytesLidos]))
}
