package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itp-backend/backend-a-co-create/helper/header"
	"github.com/itp-backend/backend-a-co-create/helper/response"
	"github.com/itp-backend/backend-a-co-create/model"
	"github.com/itp-backend/backend-a-co-create/service"
)

var (
	appJSON = "application/json"
)

func TestRouter(c *gin.Context) {
	response.BuildResponse(c, http.StatusOK, "Welcome to API Banking Innovation Co-Create", "API Test OK!")
}

func InjectFirstRole(c *gin.Context) {
	var roles = []model.Role{{Role: "admin"}, {Role: "user"}}

	contentType := header.GetContentType(c)
	var errBind error
	if contentType == appJSON {
		errBind = c.ShouldBindJSON(&roles)
	} else {
		errBind = c.ShouldBind(&roles)
	}

	if errBind != nil {
		response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", errBind.Error())
		return
	}

	anyRole := service.FindRoleForInject()

	for _, all := range roles {
		if all.Role == anyRole.Role {
			response.BuildErrResponse(c, http.StatusUnprocessableEntity, "Failed to process request", "Role has registered. This endpoint has done.")
			return
		}
	}

	createdFirstRoles := service.CreateRoles(roles)
	response.BuildResponse(c, http.StatusCreated, "Create First Roles OK!", createdFirstRoles)
}
