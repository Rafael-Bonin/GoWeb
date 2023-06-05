package main

import (
	"github.com/Rafael-Bonin/GoWeb/cmd/server/handler"
	"github.com/Rafael-Bonin/GoWeb/internal/users"
	"github.com/gin-gonic/gin"
)

func usersRouteFactory() *handler.User {
	userRepository := users.NewRepository()
	userService := users.NewService(userRepository)
	userController := handler.NewUserHandler(userService)

	return userController
}

func main() {
	userHandler := usersRouteFactory()
	router := gin.Default()
	usersRouter := router.Group("/users")

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Ola Rafael",})
	})
	usersRouter.GET("/", userHandler.GetAll())

	// usersRouter.GET("/:id", d1ginjson.GetById)

	usersRouter.POST("/create", userHandler.Create())

	router.Run()
}