package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	const_param "WebChatRoom/const"
	"WebChatRoom/tool"
)

func Login(context *gin.Context) {
	login_data := new(LoginJson)
	body := context.Request.Body
	err := tool.ParseJson(body, &login_data)
	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"reason":  "参数解析失败！",
			"data":    nil,
		})
	}

	redis_client := tool.GetRedisClient()
	defer redis_client.Close() // 每个请求创建一个链接，然后关闭。

	// result := redis_client.SAdd(const_param.RedisConst.ChatRoomKey, login_data.NickName)
	// result.val == 0 插入失败，表示已存在。1表示插入成功。 redis的set集合特性。
	val1, _ := redis_client.SAdd(const_param.RedisConst.ChatRoomKey, login_data.NickName).Result()
	if val1 == 0 {
		context.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"reason":  "昵称已存在!",
			"data":    nil,
		})
	}

	nick, token := Token(login_data.NickName)
	val2, _ := redis_client.Set(token, nick, const_param.RedisConst.Expiration).Result()

	context.SetCookie("token", token, 0, "", "", false, true)
	context.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"reason":  "成功！",
		"data":    val2, // "OK"
	})
}

type LoginJson struct {
	NickName string `json:"nick"`
}

func Token(nick string) (string, string) {
	cookie_nick := fmt.Sprintf(const_param.RedisConst.CookieTemplate, nick)
	return nick, tool.MD5(cookie_nick)
}
