package utils

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func SendErrorResponse(c *gin.Context, statusCode int, message string) {
	apiError := ErrorResponse{
		StatusCode: statusCode,
		Message:    message,
	}
	c.JSON(statusCode, apiError)
	c.Abort()
	// c.AbortWithStatus(statusCode)
	return
}

type SuccessResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func SendSuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	response := SuccessResponse{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}

	c.JSON(statusCode, response)
	return
}
