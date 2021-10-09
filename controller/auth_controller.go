package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itp-backend/backend-a-co-create/dto"
	"github.com/itp-backend/backend-a-co-create/helper/header"
	"github.com/itp-backend/backend-a-co-create/helper/response"
)

type AuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

func Login(c *gin.Context) {
	var loginDTO dto.LoginDTO
	contentType := header.GetContentType(c)

	var errDTO error
	if contentType == appJSON {
		errDTO = c.ShouldBindJSON(&loginDTO)
	} else {
		errDTO = c.ShouldBind(&loginDTO)
	}

	if errDTO != nil {
		response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", errDTO.Error())
		return
	}

}

func Register(c *gin.Context) {

}
