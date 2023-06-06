package handler

import (
	"os"
	"strconv"

	"github.com/Rafael-Bonin/GoWeb/internal/users"
	"github.com/Rafael-Bonin/GoWeb/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Nome string `json:"nome" binding:"required"`
	Sobrenome string `json:"sobrenome" binding:"required"`
	Email string `json:"email" binding:"required"`
	Idade int `json:"idade" binding:"required"`
	Altura int `json:"altura" binding:"required"`
	Ativo bool `json:"ativo" binding:"required"`
	DataDeCriacao string `json:"dataDeCriacao" binding:"required"`
}


type User struct {
	service users.Service
}

// GetAllUsers godoc
// @Summary List users
// @Tags users
// @Description get users
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /users [get]
func (u *User) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		envToken := os.Getenv("TOKEN")
		if token != envToken {
			c.JSON(401, web.NewResponse(401, nil, "usuario nao autorizado"))
			return
		}

		result, err := u.service.GetAll()

		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		c.JSON(200, web.NewResponse(200, result, ""))
		return
	}
}

// createUsers godoc
// @Summary Create users
// @Tags Users
// @Description Create users
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param user body request true "user to create"
// @Success 200 {object} web.Response
// @Router /users [post]
func (u *User) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		envToken := os.Getenv("TOKEN")
		if token != envToken {
			c.JSON(401, web.NewResponse(401, nil, "usuario nao autorizado"))
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

	result, err := u.service.Create(req.Nome, req.Sobrenome, req.Email, req.Idade, req.Altura, req.Ativo, req.DataDeCriacao)
	if err != nil {
		c.JSON(400, web.NewResponse(400, nil, err.Error()))
		return
	}

	c.JSON(200, web.NewResponse(200, result, ""))
	return 
	}
}

func (u *User) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
			token := c.GetHeader("token")
			envToken := os.Getenv("TOKEN")
			if token != envToken {
				c.JSON(401, web.NewResponse(401, nil, "usuario nao autorizado"))
				return
			}
	
			userId, e := strconv.Atoi(c.Param("id"))

			var req request
			if err := c.ShouldBindJSON(&req); err != nil || e != nil {
				c.JSON(400, web.NewResponse(400, nil, err.Error()))
				return
			}
			formatter := users.User{Id: userId, Nome: req.Nome, Sobrenome: req.Sobrenome, Email: req.Email, Idade: req.Idade, Altura: req.Altura, Ativo: req.Ativo, DataDeCriacao: req.DataDeCriacao}
			result, er := u.service.Update(formatter)

			if er != nil {
				c.JSON(400, web.NewResponse(400, nil, er.Error()))
				return
			}

			c.JSON(200, web.NewResponse(200, result, ""));
			return
	}
	
}

func (u *User) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		envToken := os.Getenv("TOKEN")
		if token != envToken {
			c.JSON(401, web.NewResponse(401, nil, "usuario nao autorizado"))
			return
		}
		userId, e := strconv.Atoi(c.Param("id"))

		if e != nil {
			c.JSON(400, web.NewResponse(400, nil, e.Error()))
			return
		}

		er := u.service.Delete(userId)
		if er != nil {
			c.JSON(404, web.NewResponse(404, nil, er.Error()))
			return
		}
		c.JSON(204, nil)
	}
}

func (u *User) SoftUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		envToken := os.Getenv("TOKEN")
		if token != envToken {
			c.JSON(401, web.NewResponse(401, nil, "usuario nao autorizado"))
			return
		}

		userId, e := strconv.Atoi(c.Param("id"))

		var req SoftUpdateRequestDTO
		if err := c.ShouldBindJSON(&req); err != nil || e != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		result, er := u.service.SoftUpdate(req.Sobrenome, req.Idade, userId)

		if er != nil {
			c.JSON(400, web.NewResponse(400, nil, er.Error()))
			return
		}

		c.JSON(200, web.NewResponse(200, result, ""));
		return
}

}

func NewUserHandler(s users.Service) *User {
	return &User{ service: s }
}