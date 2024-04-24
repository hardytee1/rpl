package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type responseSuccess struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
}

type responseError struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Errors  map[string]interface{} `json:"errors"`
}

func RespondSuccess(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, responseSuccess{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func RespondError(c *gin.Context, httpStatus int, message string, errors map[string]interface{}) {
	c.JSON(httpStatus, responseError{
		Success: false,
		Message: message,
		Errors:  errors,
	})
}
