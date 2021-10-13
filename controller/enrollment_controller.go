package controller

import (
    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "github.com/itp-backend/backend-a-co-create/helper/header"
    "github.com/itp-backend/backend-a-co-create/helper/response"
    "github.com/itp-backend/backend-a-co-create/service"
    "net/http"
    "strconv"
)

type EnrollmentController interface {
    GetEnrollmentByStatus(c *gin.Context)
    ApproveEnrollment(c *gin.Context)
}

func GetEnrollmentByStatus(c *gin.Context)  {
    status := c.Query("status")
    userData := c.MustGet("userData").(jwt.MapClaims)
    userID := userData["user_id"].(string)

    id, err := strconv.ParseUint(userID, 10, 64)
    if err != nil {
        response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", "Parsing id not working")
        return
    }
    // validate authorization admin
    getDataUser := service.FindByID(id)
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

func ApproveEnrollment(c *gin.Context)  {
    var idUserEnrollments map[string][]uint
    userData := c.MustGet("userData").(jwt.MapClaims)
    userID := userData["user_id"].(string)

    id, err := strconv.ParseUint(userID, 10, 64)
    if err != nil {
        response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", "Parsing id not working")
        return
    }
    // validate authorization admin
    getDataUser := service.FindByID(id)
    if getDataUser.RoleID != 2 {
        response.BuildErrResponse(c, http.StatusForbidden, "Failed to process request", "You're not admin")
        return
    }

    var errBind error
    contentType := header.GetContentType(c)
    if contentType == appJSON {
        errBind = c.ShouldBindJSON(&idUserEnrollments)
    } else {
        errBind = c.ShouldBind(&idUserEnrollments)
    }

    if errBind != nil {
        response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", errBind.Error())
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
