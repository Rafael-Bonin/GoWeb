package main

import (
	d1ginjson "github.com/Rafael-Bonin/GoWeb/d1GinJson"
	"github.com/gin-gonic/gin"
)



func main() {

	router := gin.Default()

	usersRouter := router.Group("/users")

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Ola Rafael",})
	})
	usersRouter.GET("/", d1ginjson.GetAll)

	usersRouter.GET("/:id", d1ginjson.GetById)

	usersRouter.POST("/create", d1ginjson.CreateProduct())

	router.Run()
}