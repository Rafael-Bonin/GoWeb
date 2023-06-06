package main

import (
	"os"

	"github.com/Rafael-Bonin/GoWeb/cmd/server/handler"
	"github.com/Rafael-Bonin/GoWeb/cmd/server/middlewares"
	"github.com/Rafael-Bonin/GoWeb/internal/users"
	"github.com/Rafael-Bonin/GoWeb/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/Rafael-Bonin/GoWeb/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go Web Meli Exercises
// @version 1.0
// @description This API Handle Go Web exercises.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	godotenv.Load()
	userHandler := usersRouteFactory()
	router := gin.Default()
	usersRouter := router.Group("/users")
	
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Ola Rafael",})
	})

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	usersRouter.Use(middlewares.AuthMiddleware()).GET("/", userHandler.GetAll())
	usersRouter.Use(middlewares.AuthMiddleware()).POST("/create", userHandler.Create())
	usersRouter.Use(middlewares.AuthMiddleware()).PUT("/update/:id", userHandler.Update())
	usersRouter.Use(middlewares.AuthMiddleware()).DELETE("/delete/:id", userHandler.Delete())
	usersRouter.Use(middlewares.AuthMiddleware()).PATCH("/softupdate/:id", userHandler.SoftUpdate())

	router.Run()
}

func usersRouteFactory() *handler.User {
	db := store.New(store.FileType, "./users.json")
	userRepository := users.NewRepository(db)
	userService := users.NewService(userRepository)
	userController := handler.NewUserHandler(userService)

	return userController
}