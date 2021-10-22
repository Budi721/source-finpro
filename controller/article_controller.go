package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/itp-backend/backend-a-co-create/dto"
	"github.com/itp-backend/backend-a-co-create/helper/response"
	"github.com/itp-backend/backend-a-co-create/helper/mc"
	"github.com/itp-backend/backend-a-co-create/service"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func CreateArticle(c *gin.Context) {
	var articleRequest *dto.Article
	if err := c.ShouldBindJSON(&articleRequest); err != nil {
		log.Error(err)
		response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", err.Error())
		return
	}

	article, err := service.CreateArticle(articleRequest)
	if err != nil {
		log.Error(err)
		response.BuildErrResponse(c, http.StatusInternalServerError, "Failed to process request", err.Error())
		return
	}
	
	response.BuildResponse(c, http.StatusCreated, "Success", article)
	return
}

func DeleteArticle(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		log.Error(err)
		response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", err.Error())
		return
	}

	err = service.DeleteArticle(id)
	result := map[string]int{
		"id_artikel_deleted": id,
	}

	if err != nil {
		log.Error(err)
		response.BuildErrResponse(c, http.StatusInternalServerError, "Failed to process request", err.Error())
		return
	}

	response.BuildResponse(c, http.StatusOK, "Success", result)
	return
}

func GetArticleById(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		log.Error(err)
		response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", err.Error())
		return
	}

	article, err := service.GetArticleById(id)
	if err != nil {
		log.Error(err)
		response.BuildErrResponse(c, http.StatusInternalServerError, "Failed to process request", err.Error())
		return
	}

	response.BuildResponse(c, http.StatusOK, "Success", article)
	return
}

func GetAllArticle(c *gin.Context) {
	articles, err := service.GetAllArticle()
	if len(articles) == 0 {
		response.BuildErrResponse(c, http.StatusNotFound, "The server has not found anything matching the Request", "Not Found")
		return
	}

	if err != nil {
		log.Error(err)
		response.BuildErrResponse(c, http.StatusInternalServerError, "Failed to process request", err.Error())
		return
	}
	
	response.BuildResponse(c, http.StatusOK, "Success", articles)
	return
}

func GetMyArticle(c *gin.Context) {
	idUser, errMC := mc.MapClaims(c)
	if errMC != nil && idUser == 0 {
		response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", errMC.Error())
		return
	}


	articles, err := service.GetArticleByIdUser(idUser)
	if err != nil {
		log.Error(err)
		response.BuildErrResponse(c, http.StatusInternalServerError, "Failed to process request", err.Error())
		return
	}

	response.BuildResponse(c, http.StatusOK, "Success", articles)
	return
}