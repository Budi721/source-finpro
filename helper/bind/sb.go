package bind

import (
	"github.com/gin-gonic/gin"
	"github.com/itp-backend/backend-a-co-create/helper/header"
)

var (
	appJSON = "application/json"
)

func SBJSON(c *gin.Context, model interface{}) (interface{}, error) {
	contentType := header.GetContentType(c)

	if contentType != appJSON {
		err := c.ShouldBindJSON(&model)
		return &model, err
	} else {
		err := c.ShouldBind(&model)
		return &model, err
	}
}
