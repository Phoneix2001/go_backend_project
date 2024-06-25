package error

import (
	"net/http"
er "banckendproject/utils/model"
	"github.com/gin-gonic/gin"
)

func WriteError(ctx *gin.Context,error string)  {
	ctx.IndentedJSON(http.StatusBadRequest, er.ErrorMessage{Status: http.StatusBadRequest, Message: error})

}