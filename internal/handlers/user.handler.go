// go-ecommerce-backend-api/internal/handlers/user.handler.go
package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/thinhnguyenwilliam/go-ecommerce-backend-api/pkg/response"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

type UserResponse struct {
	ID string `json:"id"`
}

func (uh *UserHandler) GetUserById(c *gin.Context) {
	id := c.Param("id")

	response.Success(c, UserResponse{
		ID: id,
	})
}
