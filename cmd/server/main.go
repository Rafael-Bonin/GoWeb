package main

import (
	"github.com/Rafael-Bonin/GoWeb/cmd/server/handler"
	"github.com/Rafael-Bonin/GoWeb/internal/users"
	"github.com/Rafael-Bonin/GoWeb/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func usersRouteFactory() *handler.User {
	db := store.New(store.FileType, "./users.json")
	userRepository := users.NewRepository(db)
	userService := users.NewService(userRepository)
	userController := handler.NewUserHandler(userService)

	return userController
}

func main() {
	godotenv.Load()
	userHandler := usersRouteFactory()
	router := gin.Default()
	usersRouter := router.Group("/users")

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Ola Rafael",})
	})

	usersRouter.GET("/", userHandler.GetAll())

	usersRouter.POST("/create", userHandler.Create())

	usersRouter.PUT("/update/:id", userHandler.Update())

	usersRouter.DELETE("/delete/:id", userHandler.Delete())

	usersRouter.PATCH("/softupdate/:id", userHandler.SoftUpdate())

	router.Run()
}