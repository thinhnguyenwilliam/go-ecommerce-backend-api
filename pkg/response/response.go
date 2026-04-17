// go-ecommerce-backend-api/pkg/response/response.go
package response

import "github.com/gin-gonic/gin"

type ResponseData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// Success
func SuccessResponse(c *gin.Context, code int, data any) {
	c.JSON(200, ResponseData{
		Code:    code,
		Message: GetMsg(code),
		Data:    data,
	})
}

func Success(c *gin.Context, data any) {
	SuccessResponse(c, ErrorCodeSuccess, data)
}

// Error
func ErrorResponse(c *gin.Context, code int) {
	c.JSON(200, ResponseData{
		Code:    code,
		Message: GetMsg(code),
		Data:    nil,
	})
}
