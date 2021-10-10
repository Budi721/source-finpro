package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/itp-backend/backend-a-co-create/dto"
	"github.com/itp-backend/backend-a-co-create/helper/header"
	"github.com/itp-backend/backend-a-co-create/helper/response"
	"github.com/itp-backend/backend-a-co-create/model"
	"github.com/itp-backend/backend-a-co-create/service"
)

type AuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	Logout()
}

func Login(c *gin.Context) {
	var loginDTO dto.LoginDTO
	contentType := header.GetContentType(c)

	var errDTO error
	if contentType == appJSON {
		errDTO = c.ShouldBindJSON(&loginDTO)
	} else {
		errDTO = c.ShouldBind(&loginDTO)
	}

	if errDTO != nil {
		response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", errDTO.Error())
		return
	}

	authResult := service.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if val, ok := authResult.(model.User); ok {
		generatedToken := service.GenerateToken(strconv.FormatUint(val.ID, 10))
		val.Token = generatedToken
		response.BuildResponse(c, http.StatusOK, "Login OK!", val)
		return
	}
	response.BuildErrResponse(c, http.StatusUnauthorized, "Please check your credential", "Invalid credential")
}

func Register(c *gin.Context) {
	var registerDTO dto.RegisterDTO
	contentType := header.GetContentType(c)

	var errDTO error
	if contentType == appJSON {
		errDTO = c.ShouldBindJSON(&registerDTO)
	} else {
		errDTO = c.ShouldBind(&registerDTO)
	}

	if errDTO != nil {
		response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", errDTO.Error())
		return
	}

	if service.IsDuplicateEmail(registerDTO.Email) {
		response.BuildErrResponse(c, http.StatusConflict, "Failed to process request", "Duplicate email")
		return
	}

	createdUser := service.CreateUser(registerDTO)
	generatedToken := service.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
	createdUser.Token = generatedToken
	response.BuildResponse(c, http.StatusCreated, "Register OK!", createdUser)
}
