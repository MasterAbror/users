package environment

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// GinRouter -> Gin Router
type GinRouter struct {
	Gin *gin.Engine
}

// NewGinRouter all the routes are defined here
func NewGinRouter() GinRouter {

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"https://edu.jakarta.go.id"}
	corsConfig.AllowCredentials = true

	httpRouter := gin.Default()
	httpRouter.Use(cors.New(corsConfig))

	httpRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello Abror... Users API is up and running..."})
	})
	return GinRouter{
		Gin: httpRouter,
	}

}
