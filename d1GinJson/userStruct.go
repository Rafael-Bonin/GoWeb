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
	DataDeCriacao string `json:"dataDeCriacao"`
}

var AllUsers []User = []User{}
var id int = 0

func CreateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("token")

		if token != "tokentest123" {
			c.JSON(401, gin.H{ "error": "usuario nao autorizado" })
			return
		}

		var req User
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		}
		if req.Altura == 0 { c.JSON(400, gin.H{ "error": "campo altura nao preenchido" })
	return }
		if req.Nome == "" { c.JSON(400, gin.H{ "error": "campo nome nao preenchido" }) 
	return}
		if req.Sobrenome == "" { c.JSON(400, gin.H{ "error": "campo sobrenome nao preenchido" })
	return}
		if req.Email == "" {c.JSON(400, gin.H{ "error": "campo email nao preenchido" })
	return}
		if req.Idade == 0 {c.JSON(400, gin.H{ "error": "campo idade nao preenchido" })
	return}
		if req.Ativo == false {c.JSON(400, gin.H{ "error": "campo ativo nao preenchido" })
	return}
		if req.DataDeCriacao == "" {c.JSON(400, gin.H{ "error": "campo dataDeCriacao nao preenchido" })
	return}


		id += 1
		req.Id = id 
		AllUsers = append(AllUsers, req)
		c.JSON(200, req);
		return
	}
}

func GetAll(c *gin.Context) {

	c.JSON(200, gin.H{ "users": AllUsers })
}

func GetById(c *gin.Context) {

	var user User

	for _, value := range AllUsers {
		if fmt.Sprint(value.Id) == c.Param("id") {
			user = value
			}
	}
	if user.Id == 0 {
		c.JSON(404, gin.H{"error": "no user was found"})
	} else {
		c.JSON(200, gin.H{ "User": user })
	}
	
}