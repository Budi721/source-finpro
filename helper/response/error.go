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
	splittedErr := strings.Split(err, "\n")
	res := ErrResponse{
		Status:     false,
		StatusCode: sc,
		Message:    message,
		Error:      splittedErr,
	}

	c.AbortWithStatusJSON(sc, res)
}
