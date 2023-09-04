package utilities

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSONResponse represents a consistent JSON response structure
type JSONResponse struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// SuccessResponse sends a success JSON response with data
func SuccessResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, JSONResponse{
		Message: message,
		Data:    data,
	})
}

// ErrorResponse sends an error JSON response with an error message
func ErrorResponse(c *gin.Context, statusCode int, errorMessage string) {
	c.JSON(statusCode, JSONResponse{
		Error: errorMessage,
	})
}
