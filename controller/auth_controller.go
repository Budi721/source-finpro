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
		if val.RoleID != uint(loginDTO.LoginAs) {
			response.BuildErrResponse(c, http.StatusForbidden, "Failed to process request", "LoginAs not match. You're login not for this endpoint.")
			return
		}

		role := service.FindRoleID(uint64(val.RoleID))

		generatedToken := service.GenerateToken(strconv.FormatUint(val.ID, 10))
		res := dto.ResponseLogRegDTO{
			ID:    val.ID,
			Name:  val.Name,
			Role:  role.Role,
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

	getRole := service.FindRoleID(registerDTO.RoleID)
	if getRole.ID == 0 {
		response.BuildErrResponse(c, http.StatusNotFound, "Failed to process request", "Role not found")
		return
	}

	createdUser := service.CreateUser(registerDTO)
	role := service.FindRoleID(registerDTO.RoleID)
	generatedToken := service.GenerateToken(strconv.FormatUint(createdUser.IdUser, 10))
	res := dto.ResponseLogRegDTO{
		ID:            createdUser.IdUser,
		Name:          createdUser.NamaLengkap,
		Email:         createdUser.Email,
		Role:          role.Role,
		Token:         generatedToken,
		TopikDiminati: createdUser.TopikDiminati,
	}
	response.BuildResponse(c, http.StatusCreated, "Register OK!", res)
}
