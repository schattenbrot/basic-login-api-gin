package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// writeJSON is the helper function for sending back an HTTP response.
func WriteJSON(c *gin.Context, status int, data ...interface{}) {
	c.Header("Content-Type", "application/json")

	if len(data) == 0 {
		c.JSON(status, "")
		return
	}

	c.JSON(status, data[0]) // data is the first argument
}

// errorJSON is the helper function for creating an error message.
// This internally then runs the writeJSON function to send the HTTP response.
func ErrorJSON(c *gin.Context, err error, status ...int) {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	type jsonError struct {
		StatusCode int    `json:"statusCode"`
		Message    string `json:"message"`
	}

	theError := jsonError{
		StatusCode: statusCode,
		Message:    err.Error(),
	}

	c.Header("Content-Type", "application/json")
	c.JSON(statusCode, theError)
}
