package middleware

import (
	"log"
	"net/http"
	"strings"
	"testcocreate/helper/response"
	"testcocreate/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.BuildErrResponse(c, http.StatusNotFound, "Failed to process request", "No token found")
			return
		}

		bearer := strings.HasPrefix(authHeader, "Bearer")
		if !bearer {
			response.BuildErrResponse(c, http.StatusUnauthorized, "Failed to process request", "Bearer token rules")
			return
		}

		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[user_id]:", claims["user_id"])
			log.Println("Claims[issuer]:", claims["issuer"])
		} else {
			log.Println(err)
			response.BuildErrResponse(c, http.StatusUnauthorized, "Token is not valid", err.Error())
			return
		}
	}
}
