package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	const_param "WebChatRoom/const"
	"WebChatRoom/tool"
)

func PostMessage(context *gin.Context) {

	token, err := context.Cookie("token")
	if err != nil {
		context.Redirect(http.StatusTemporaryRedirect, "/")
	}

	redis_client := tool.GetRedisClient()
	defer redis_client.Close() // 每个请求创建一个链接，然后关闭。

	nick, err := redis_client.Get(token).Result()
	if err != nil {
		context.Redirect(http.StatusTemporaryRedirect, "/")
	}

	msg_data := new(MessageJson)
	body := context.Request.Body
	err = tool.ParseJson(body, &msg_data)
	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"reason":  "参数解析失败!",
			"data":    nil,
		})
	}
	// msg_data.PostTime = time.Now().Format("YYYY-MM-DD hh:mm:ss") 奇葩：这里时间格式用go出生时间
	msg_data.PostTime = time.Now().Format("2006-01-02 15:04:05")
	msg_data.NickName = nick

	msg_store := msg_data.GetMessageStoreString()
	val, _ := redis_client.RPush(const_param.RedisConst.ChatListKey, msg_store).Result()
	if val != 0 {
		context.JSON(http.StatusOK, map[string]interface{}{
			"success": true,
			"reason":  "OK!",
			"data":    msg_data,
		})
	}
}

type MessageJson struct {
	NickName string `json:"nick"`
	Message  string `json:"msg"`
	PostTime string `json:"post_time"`
}

func (ms MessageJson) GetMessageStoreString() []byte {
	res, err := json.Marshal(ms)
	if err != nil {
		panic(err)
	}
	return res
}

func BackMessageJson(ms string) *MessageJson {
	mj := new(MessageJson)
	err := json.Unmarshal([]byte(ms), mj)
	if err != nil {
		panic(err)
	}
	return mj
}
