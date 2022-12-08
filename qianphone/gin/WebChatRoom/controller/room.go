package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Room(context *gin.Context) {

	_, err := context.Cookie("token")
	if err != nil {
		context.Redirect(http.StatusTemporaryRedirect, "/")
	}

	context.HTML(http.StatusOK, "chatroom.html", gin.H{
		"title": "Main website",
	})
}
