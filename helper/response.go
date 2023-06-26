package helper

import "github.com/gin-gonic/gin"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessJSON(ctx *gin.Context, message string, status int, data interface{}) Response {
	resp := Response{
		Success: true,
		Message: message,
		Data:    data,
	}

	return resp
}

func ErrorJSON(ctx *gin.Context, message string, status int, data interface{}) Response {
	resp := Response{
		Success: false,
		Message: message,
		Data:    data,
	}

	return resp
}
