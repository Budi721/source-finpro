package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/itp-backend/backend-a-co-create/dto"
	"github.com/itp-backend/backend-a-co-create/helper/header"
	"github.com/itp-backend/backend-a-co-create/helper/mc"
	"github.com/itp-backend/backend-a-co-create/helper/response"
	"github.com/itp-backend/backend-a-co-create/service"
)

type RoleController interface {
	GetAllRoles(c *gin.Context)
	CreateRole(c *gin.Context)
	MyRole(c *gin.Context)
	UpdateRole(c *gin.Context)
	DeleteRole(c *gin.Context)
}

func GetAllRoles(c *gin.Context) {
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

	allRoles := service.GetAllRole()
	response.BuildResponse(c, http.StatusOK, "All Data OK!", allRoles)
}

func CreateRole(c *gin.Context) {
	var role dto.RoleDTO

	contentType := header.GetContentType(c)
	var errBind error
	if contentType == appJSON {
		errBind = c.ShouldBindJSON(&role)
	} else {
		errBind = c.ShouldBind(&role)
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

	getDataUser := service.FindByID(userid)
	if getDataUser.RoleID != 1 {
		response.BuildErrResponse(c, http.StatusForbidden, "Failed to process request", "You're not admin")
		return
	}

	createdRole := service.CreateRole(role)
	response.BuildResponse(c, http.StatusCreated, "Create Role OK!", createdRole)
}

func MyRole(c *gin.Context) {
	userid, errMC := mc.MapClaims(c)
	if errMC != nil && userid == 0 {
		response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", errMC.Error())
		return
	}

	getUser := service.FindByID(userid)
	getRole := service.FindRoleID(uint64(getUser.RoleID))
	response.BuildResponse(c, http.StatusOK, "Get Role OK!", getRole)
}

func DeleteRole(c *gin.Context) {
	roleID := c.Param("id")

	roleid, errPR := strconv.ParseUint(roleID, 10, 64)
	if errPR != nil {
		response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", errPR.Error())
		return
	}

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

	roleDelete := service.DeleteRole(roleid)
	response.BuildResponse(c, http.StatusAccepted, "Role has deleted", roleDelete)
}
