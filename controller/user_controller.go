package controller

import "github.com/gin-gonic/gin"

type UserController interface {
	GetAllUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	Profile(c *gin.Context)
}
