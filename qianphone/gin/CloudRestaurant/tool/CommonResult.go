package tool

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS int = 200
	FAILURE int = 401
)

func Success(context *gin.Context, v interface{}) {
	context.JSON(http.StatusOK, map[string]interface{}{
		"code": SUCCESS,
		"msg":  "成功！",
		"data": v,
	})
}

func Failure(context *gin.Context, v interface{}) {
	context.JSON(http.StatusBadRequest, map[string]interface{}{
		"code": FAILURE,
		"msg":  v,
		"data": nil,
	})
}
