package controller

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/itp-backend/backend-a-co-create/dto"
	"github.com/itp-backend/backend-a-co-create/helper/header"
	"github.com/itp-backend/backend-a-co-create/helper/response"
	"github.com/itp-backend/backend-a-co-create/service"
)

type UserController interface {
	GetAllUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	ChangePassword(c *gin.Context)
	MyProfile(c *gin.Context)
}

func GetAllUser(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)

	userID := userData["user_id"].(string)
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", "Parsing id not working")
		return
	}

	getDataUser := service.FindByID(id)
	if getDataUser.RoleID != 1 {
		response.BuildErrResponse(c, http.StatusForbidden, "Failed to process request", "You're not admin")
		return
	}

	allUsers := service.GetAllUser()
	response.BuildResponse(c, http.StatusOK, "All Data OK!", allUsers)
}

func UpdateUser(c *gin.Context) {
	var user dto.UserUpdateDTO
	contentType := header.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)

	var errBind error
	if contentType == appJSON {
		errBind = c.ShouldBindJSON(&user)
	} else {
		errBind = c.ShouldBind(&user)
	}

	if errBind != nil {
		response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", errBind.Error())
		return
	}

	userID := userData["user_id"].(string)
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", "Parsing id not working")
		return
	}

	userUpdate := service.UpdateUser(id, user)
	response.BuildResponse(c, http.StatusCreated, "Update Data OK!", userUpdate)
}

func ChangePassword(c *gin.Context) {
	var userPass dto.ChangePasswordDTO
	contentType := header.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)

	var errBind error
	if contentType == appJSON {
		errBind = c.ShouldBindJSON(&userPass)
	} else {
		errBind = c.ShouldBind(&userPass)
	}

	if errBind != nil {
		response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", errBind.Error())
		return
	}

	if userPass.Password != userPass.ConfirmPassword {
		response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", "Password not matched")
		return
	}

	userID := userData["user_id"].(string)
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", "Parsing id not working")
		return
	}

	userChangePass := service.ChangePassword(id, userPass)
	response.BuildResponse(c, http.StatusCreated, "Change Password OK!", userChangePass)
}

func MyProfile(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)

	userID := userData["user_id"].(string)
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", "Parsing id not working")
		return
	}

	getUser := service.FindByID(id)
	response.BuildResponse(c, http.StatusOK, "Get Profile OK!", getUser)
}
