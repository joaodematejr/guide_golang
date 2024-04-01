package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Definindo a estrutura do modelo de usuário
type User struct {
	ID   uint `gorm:"primary_key"`
	Name string
	Age  uint
}

func main() {
	// Conectando ao banco de dados SQLite
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Falha ao conectar ao banco de dados!")
	}
	defer db.Close()

	// Automatizando a migração do modelo, criando a tabela 'users'
	db.AutoMigrate(&User{})

	// Criando um novo usuário
	user := User{Name: "John Doe", Age: 30}
	db.Create(&user)
	fmt.Println("Novo usuário criado:", user)

	// Buscando um usuário pelo ID
	var retrievedUser User
	db.First(&retrievedUser, user.ID)
	fmt.Println("Usuário recuperado:", retrievedUser)

	// Atualizando os dados do usuário
	db.Model(&retrievedUser).Update("Age", 35)
	fmt.Println("Usuário atualizado:", retrievedUser)

	// Deletando um usuário
	db.Delete(&retrievedUser)
	fmt.Println("Usuário deletado:", retrievedUser)
}
