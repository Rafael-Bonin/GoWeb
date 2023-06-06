package middlewares

import (
	"os"

	"github.com/gin-gonic/gin"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}


func AuthMiddleware() gin.HandlerFunc {
	envToken := os.Getenv("TOKEN")

	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != envToken {
			respondWithError(c, 401, "usuario não autorizado, por favor faça login")
			return
		}
		c.Next()
	}
}