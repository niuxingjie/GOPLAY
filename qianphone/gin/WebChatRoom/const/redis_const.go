package const_param

import (
	"time"
)

type redisConst struct {
	ChatRoomKey    string
	CookieTemplate string
	Expiration     time.Duration
	ChatListKey    string
}

var RedisConst = redisConst{ChatRoomKey: "chat_room_key",
	CookieTemplate: "cookie:%s",
	Expiration:     60 * 15 * time.Second,
	ChatListKey:    "chst_list",
}
