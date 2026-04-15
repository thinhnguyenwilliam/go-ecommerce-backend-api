// go-ecommerce-backend-api/internal/routers/router.go
package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/thinhnguyenwilliam/go-ecommerce-backend-api/internal/handlers"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	userHandler := handlers.NewUserHandler()

	v1 := r.Group("/v1/2026")
	{
		v1.GET("/users/:id", userHandler.GetUserById)
	}

	return r
}
