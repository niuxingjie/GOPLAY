package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	const_param "WebChatRoom/const"
	"WebChatRoom/tool"
)

func GetChatList(context *gin.Context) {
	redis_client := tool.GetRedisClient()
	defer redis_client.Close() // 每个请求创建一个链接，然后关闭。

	val, _ := redis_client.LRange(const_param.RedisConst.ChatListKey, -20, -1).Result()
	var msg_slice []MessageJson = make([]MessageJson, 0, 20)
	for _, msg := range val {
		msg_slice = append(msg_slice, *BackMessageJson(msg))
	}
	msg_string, _ := json.Marshal(msg_slice)
	context.String(http.StatusOK, string(msg_string))
}
