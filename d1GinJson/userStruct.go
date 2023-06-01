package d1ginjson

import (
	"fmt"

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


func GetAll(c *gin.Context) {
	mockedUsers := []User{{Id: 1, Nome: "Rafael", Sobrenome: "Bonin", Email: "rafalyf11@gmail.com", Idade: 21, Altura: 179, Ativo: true, DataDeCriacao: "22/05/2023"}, {Id: 2, Nome: "leafar", Sobrenome: "ninob", Email: "fake@email.com", Idade: 30, Altura: 160, Ativo: true, DataDeCriacao: "11/01/2002"}}
	var filtered  []User
	for _, value := range mockedUsers {
		if value.Nome == c.Query("nome") {
		filtered = append(filtered, value)
		}
	}

	c.JSON(200, gin.H{ "users": filtered })
}

func GetById(c *gin.Context) {
	mockedUsers := []User{{Id: 1, Nome: "Rafael", Sobrenome: "Bonin", Email: "rafalyf11@gmail.com", Idade: 21, Altura: 179, Ativo: true, DataDeCriacao: "22/05/2023"}, {Id: 2, Nome: "leafar", Sobrenome: "ninob", Email: "fake@email.com", Idade: 30, Altura: 160, Ativo: true, DataDeCriacao: "11/01/2002"}}

	var user User

	for _, value := range mockedUsers {
		if fmt.Sprint(value.Id) == c.Param("id") {
			user = value
			}
	}

	c.JSON(200, gin.H{ "User": user })
}