package response

import "github.com/gin-gonic/gin"

type Response struct {
	Status     bool        `json:"status"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Result     interface{} `json:"results"`
}

func BuildResponse(c *gin.Context, sc int, message string, result interface{}) {
	res := Response{
		Status:     true,
		StatusCode: sc,
		Message:    message,
		Result:     result,
	}

	c.JSON(sc, res)
}
