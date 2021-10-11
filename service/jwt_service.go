package service

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/itp-backend/backend-a-co-create/config"
	log "github.com/sirupsen/logrus"
)

type JWTService interface {
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

var (
	issuer = "ydhnwb"
	secret = config.Init().JWTSecret
)

func GenerateToken(userID string) string {
	claims := &jwtCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 3).Unix(),
			Issuer:    issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := parseToken.SignedString([]byte(secret))
	if err != nil {
		log.Warning("Signed Token Error")
	}

	return signedToken
}

func ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})
}
