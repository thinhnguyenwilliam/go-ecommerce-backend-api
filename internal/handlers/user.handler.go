// go-ecommerce-backend-api/internal/handlers/user.handler.go
package handlers

import "github.com/gin-gonic/gin"

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (uh *UserHandler) GetUserById(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, gin.H{
		"message": "get user by id",
		"id":      id,
	})
}
