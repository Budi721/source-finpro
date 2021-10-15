package mc

import (
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func MapClaims(c *gin.Context) (uint64, error) {
	userData := c.MustGet("userData").(jwt.MapClaims)

	userID := userData["user_id"].(string)

	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return 0, err
	}

	return id, nil
}
