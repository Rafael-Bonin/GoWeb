package handler

import (
	"github.com/Rafael-Bonin/GoWeb/internal/users"
	"github.com/gin-gonic/gin"
)

type request struct {
	Nome string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Email string `json:"email"`
	Idade int `json:"idade"`
	Altura int `json:"altura"`
	Ativo bool `json:"ativo"`
	DataDeCriacao string `json:"dataDeCriacao"`
}

type User struct {
	service users.Service
}

func (u *User) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token != "tokentest123" {
			c.JSON(401, gin.H{ "error": "usuario nao autorizado" })
			return
		}

		result, err := u.service.GetAll()

		if err != nil {
			c.JSON(400, gin.H{ "error": "erro inesperado" })
			return
		}

		c.JSON(200, result)
		return
	}
}

func (u *User) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token != "tokentest123" {
			c.JSON(401, gin.H{ "error": "usuario nao autorizado" })
			return
		}

		var req request
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

	result, err := u.service.Create(req.Nome, req.Sobrenome, req.Email, req.Idade, req.Altura, req.Ativo, req.DataDeCriacao)
	if err != nil {
		c.JSON(400, gin.H{ "error": "erro inesperado" })
	}

	c.JSON(200, result)

	}
}

func (u *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func NewUserHandler(s users.Service) *User {
	return &User{ service: s }
}