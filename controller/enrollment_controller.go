package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itp-backend/backend-a-co-create/helper/mc"
	"github.com/itp-backend/backend-a-co-create/helper/response"
	"github.com/itp-backend/backend-a-co-create/service"
)

type EnrollmentController interface {
	GetEnrollmentByStatus(c *gin.Context)
	ApproveEnrollment(c *gin.Context)
}

func GetEnrollmentByStatus(c *gin.Context) {
	status := c.Query("status")
	userid, errMC := mc.MapClaims(c)
	if errMC != nil && userid == 0 {
		response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", errMC.Error())
		return
	}

	// validate authorization admin
	getDataUser := service.FindByID(userid)
	if getDataUser.RoleID != 2 {
		response.BuildErrResponse(c, http.StatusForbidden, "Failed to process request", "You're not admin")
		return
	}

	enrollments, err := service.GetEnrollmentByStatus(status)
	if len(enrollments) == 0 {
		response.BuildErrResponse(c, http.StatusNotFound, "The server has not found anything matching the Request", "Not Found")
		return
	}

	if err != nil {
		response.BuildErrResponse(c, http.StatusInternalServerError, "Failed to process request", err.Error())
		return
	}
	response.BuildResponse(c, http.StatusOK, "All Data OK!", enrollments)
}

func ApproveEnrollment(c *gin.Context) {
	var idUserEnrollments map[string][]uint
	userid, errMC := mc.MapClaims(c)
	if errMC != nil && userid == 0 {
		response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", errMC.Error())
		return
	}
	// validate authorization admin
	getDataUser := service.FindByID(userid)
	if getDataUser.RoleID != 2 {
		response.BuildErrResponse(c, http.StatusForbidden, "Failed to process request", "You're not admin")
		return
	}

	if err := c.ShouldBindJSON(&idUserEnrollments); err != nil {
		response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", err.Error())
		return
	}

	enrollments, err := service.ApproveEnrollment(idUserEnrollments["user_ids"])
	if len(enrollments) == 0 {
		response.BuildErrResponse(c, http.StatusNotFound, "The server has not found anything matching the Request", "Not Found")
		return
	}

	if err != nil {
		response.BuildErrResponse(c, http.StatusInternalServerError, "Failed to process request", err.Error())
		return
	}
	response.BuildResponse(c, http.StatusOK, "Enrollment Approved!", enrollments)
}
