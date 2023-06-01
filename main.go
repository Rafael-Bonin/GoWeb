package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Id int `json:"id"`
	Nome string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Email string `json:"email"`
	Idade int `json:"idade"`
	Altura int `json:"altura"`
	Ativo bool `json:"ativo"`
	DataDeCriacao string `json:"data_de_criacao"`
}

func main() {
	mockedUsers := []User{{Id: 1, Nome: "Rafael", Sobrenome: "Bonin", Email: "rafalyf11@gmail.com", Idade: 21, Altura: 179, Ativo: true, DataDeCriacao: "22/05/2023"}, {Id: 2, Nome: "leafar", Sobrenome: "ninob", Email: "fake@email.com", Idade: 30, Altura: 160, Ativo: true, DataDeCriacao: "11/01/2002"}}
	router := gin.Default()

	router.GET("ola-eu", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Ola Rafael",})
	})
	router.GET("users", func(c *gin.Context) {
		c.JSON(200, gin.H{ "users": mockedUsers })
	})

	router.Run()
}