package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itp-backend/backend-a-co-create/dto"
	"github.com/itp-backend/backend-a-co-create/helper/header"
	"github.com/itp-backend/backend-a-co-create/helper/mc"
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
	userid, errMC := mc.MapClaims(c)
	if errMC != nil && userid == 0 {
		response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", errMC.Error())
		return
	}

	getDataUser := service.FindByID(userid)
	if getDataUser.RoleID != 1 {
		response.BuildErrResponse(c, http.StatusForbidden, "Failed to process request", "You're not admin")
		return
	}

	allUsers := service.GetAllUser()
	response.BuildResponse(c, http.StatusOK, "All Data OK!", allUsers)
}

func GetAllJustUser(c *gin.Context) {
	userid, errMC := mc.MapClaims(c)
	if errMC != nil && userid == 0 {
		response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", errMC.Error())
		return
	}

	role := service.FindRoleID(2)
	justAllUser := service.GetAllJustUser()
	for _, user := range justAllUser {
		user.Role = role.Role
	}

	response.BuildResponse(c, http.StatusOK, "All just user OK!", justAllUser)
}

func UpdateUser(c *gin.Context) {
	var user dto.UserUpdateDTO

	contentType := header.GetContentType(c)
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

	userid, errMC := mc.MapClaims(c)
	if errMC != nil && userid == 0 {
		response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", errMC.Error())
		return
	}

	userUpdate := service.UpdateUser(userid, user)
	response.BuildResponse(c, http.StatusCreated, "Update Data OK!", userUpdate)
}

func ChangePassword(c *gin.Context) {
	var userPass dto.ChangePasswordDTO

	contentType := header.GetContentType(c)
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

	userid, errMC := mc.MapClaims(c)
	if errMC != nil && userid == 0 {
		response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", errMC.Error())
		return
	}

	userChangePass := service.ChangePassword(userid, userPass)
	response.BuildResponse(c, http.StatusCreated, "Change Password OK!", userChangePass)
}

func MyProfile(c *gin.Context) {
	userid, errMC := mc.MapClaims(c)
	if errMC != nil && userid == 0 {
		response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", errMC.Error())
		return
	}

	getUser := service.FindByID(userid)
	response.BuildResponse(c, http.StatusOK, "Get Profile OK!", getUser)
}
