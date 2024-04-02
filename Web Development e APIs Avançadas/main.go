package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Struct que representa um usuário
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users []User // Simula um banco de dados em memória

func main() {
	r := gin.Default()

	// Endpoint para obter todos os usuários
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, users)
	})

	// Endpoint para obter um usuário pelo ID
	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, user := range users {
			if user.ID == id {
				c.JSON(http.StatusOK, user)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	})

	// Endpoint para criar um novo usuário
	r.POST("/users", func(c *gin.Context) {
		var newUser User
		if err := c.BindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		users = append(users, newUser)
		c.Status(http.StatusCreated)
	})

	// Endpoint para atualizar um usuário existente
	r.PUT("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updatedUser User
		if err := c.BindJSON(&updatedUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for i, user := range users {
			if user.ID == id {
				users[i] = updatedUser
				c.Status(http.StatusOK)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	})

	// Endpoint para deletar um usuário
	r.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, user := range users {
			if user.ID == id {
				users = append(users[:i], users[i+1:]...)
				c.Status(http.StatusOK)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	})

	// Executa o servidor na porta 8080
	r.Run(":8080")
}
