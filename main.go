package main

import (
	d1ginjson "github.com/Rafael-Bonin/GoWeb/d1GinJson"
	"github.com/gin-gonic/gin"
)



func main() {

	router := gin.Default()

	router.GET("ola-eu", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Ola Rafael",})
	})
	router.GET("users", d1ginjson.GetAll)

	router.GET("users/:id", d1ginjson.GetById)

	router.Run()
}