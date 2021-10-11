package response

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type ErrResponse struct {
	Status     bool        `json:"status"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Error      interface{} `json:"errors"`
}

func BuildErrResponse(c *gin.Context, sc int, message string, err string) {
	hasContains := strings.Contains(err, "\n")
	var errString interface{}
	if hasContains {
		errString = strings.Split(err, "\n")
	} else {
		errString = err
	}

	res := ErrResponse{
		Status:     false,
		StatusCode: sc,
		Message:    message,
		Error:      errString,
	}

	c.AbortWithStatusJSON(sc, res)
}
