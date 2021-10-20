package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/itp-backend/backend-a-co-create/dto"
	"github.com/itp-backend/backend-a-co-create/helper/mc"
	"github.com/itp-backend/backend-a-co-create/helper/response"
	"github.com/itp-backend/backend-a-co-create/service"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

const layoutFormatDate = "2006-01-02"


func CreateProject(c *gin.Context) {

	userid, errMC := mc.MapClaims(c)
	if errMC != nil && userid == 0 {
		response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", errMC.Error())
		return
	}


	var project dto.Project
	project.Creator = userid
	if err := c.ShouldBindJSON(&project); err != nil {
		log.Error(err)
		response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", err.Error())
		return
	}

	date, errDate := time.Parse(layoutFormatDate, project.StartDate)
	if errDate != nil {
		response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", errDate.Error())
		return
	}

	project.Date = date

	createdProject, err := service.CreateProject(&project)
	if err != nil {
		log.Error(err)
		response.BuildErrResponse(c, http.StatusInternalServerError, "Failed to process request", err.Error())
		return
	}

	response.BuildResponse(c, http.StatusCreated, "Success created project", createdProject)
	return
}

func DetailProject(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		log.Error(err)
		response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", err.Error())
		return
	}
	project, err := service.GetDetailProject(id)
	if err != nil {
		log.Error(err)
		response.BuildErrResponse(c, http.StatusInternalServerError, "Failed to process request", err.Error())
		return
	}

	response.BuildResponse(c, http.StatusOK, "Success !", project)
}

func DeleteProject(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		log.Error(err)
		response.BuildErrResponse(c, http.StatusInternalServerError, "Failed to process request", err.Error())
		return
	}

	err = service.DeleteProject(id)
	result := map[string]int{
		"id_project_deleted": id,
	}

	if err != nil {
		log.Error(err)
		response.BuildErrResponse(c, http.StatusInternalServerError, "Failed to process request", err.Error())
		return
	}

	response.BuildResponse(c, http.StatusOK, "Success deleted project!!", result)
}

func ProjectByInvitedUserId(c *gin.Context) {
	invitedUser := c.Query("invited_user_id")
	id, err := strconv.Atoi(invitedUser)
	if err != nil {
		log.Error(err)
		response.BuildErrResponse(c, http.StatusInternalServerError, "Failed to process request", err.Error())
		return
	}

	project, err := service.GetProjectByInvitedUser(id)
	if err != nil {
		log.Error(err)
		response.BuildErrResponse(c, http.StatusInternalServerError, "Failed to process request", err.Error())
		return
	}

	response.BuildResponse(c, http.StatusInternalServerError, "Failed to process request", project)
}

func AcceptProject(c *gin.Context) {
	var projectInvitation dto.ProjectInvitation

	if err := c.ShouldBindJSON(&projectInvitation); err != nil {
		log.Error(err)
		response.BuildErrResponse(c, http.StatusInternalServerError, "Failed to process request", err.Error())
		return
	}

	projectUpdated, err := service.UpdateInvitation(projectInvitation)
	if err != nil {
		log.Error(err)
		response.BuildErrResponse(c, http.StatusInternalServerError, "Failed to process request", err.Error())
		return
	}

	response.BuildResponse(c, http.StatusOK, "Failed to process request", projectUpdated)
	return
}

func GetAllProject(c *gin.Context) {
	projects, err := service.GetAllProject()
	if len(projects) == 0 {
		response.BuildErrResponse(c, http.StatusNotFound, "The server has not found anything matching the Request", "Not Found")
		return
	}
	
	if err != nil {
		log.Error(err)
		response.BuildErrResponse(c, http.StatusInternalServerError, "Failed to process request", err.Error())
		return
	}
	
	response.BuildResponse(c, http.StatusOK, "Success", projects)
	return
}
