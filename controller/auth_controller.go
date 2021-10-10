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

	var errBind error
	if contentType == appJSON {
		errBind = c.ShouldBindJSON(&loginDTO)
	} else {
		errBind = c.ShouldBind(&loginDTO)
	}

	if errBind != nil {
		response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", errBind.Error())
		return
	}

	authResult := service.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if val, ok := authResult.(model.User); ok {
		generatedToken := service.GenerateToken(strconv.FormatUint(val.ID, 10))
		res := dto.ResponseLogRegDTO{
			ID:    val.ID,
			Name:  val.Name,
			Email: val.Email,
			Token: generatedToken,
		}
		response.BuildResponse(c, http.StatusOK, "Login OK!", res)
		return
	}
	response.BuildErrResponse(c, http.StatusUnauthorized, "Please check your credential", "Invalid credential")
}

func Register(c *gin.Context) {
	var registerDTO dto.RegisterDTO
	contentType := header.GetContentType(c)

	var errBind error
	if contentType == appJSON {
		errBind = c.ShouldBindJSON(&registerDTO)
	} else {
		errBind = c.ShouldBind(&registerDTO)
	}

	if errBind != nil {
		response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", errBind.Error())
		return
	}

	if service.IsDuplicateEmail(registerDTO.Email) {
		response.BuildErrResponse(c, http.StatusConflict, "Failed to process request", "Duplicate email")
		return
	}

	createdUser := service.CreateUser(registerDTO)
	generatedToken := service.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
	res := dto.ResponseLogRegDTO{
		ID:    createdUser.ID,
		Name:  createdUser.Name,
		Email: createdUser.Email,
		Token: generatedToken,
	}
	response.BuildResponse(c, http.StatusCreated, "Register OK!", res)
}
