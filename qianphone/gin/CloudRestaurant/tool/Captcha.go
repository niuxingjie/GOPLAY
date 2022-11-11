package tool

import (
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/mojocn/base64Captcha"
)

//configJsonBody json request body.
type ConfigJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

type RedisStore struct {
	Client *redis.Client
}

func (rc RedisStore) Set(id string, value string) error {
	err := rc.Client.Set(id, value, time.Minute*10).Err()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (rc RedisStore) Get(id string, clear bool) string {
	value, err := rc.Client.Get(id).Result()
	if err != nil {
		log.Println(err)
		return ""
	}

	if clear {
		err := rc.Client.Del(id).Err()
		if err != nil {
			log.Println(err)
			return ""
		}
	}
	return value
}

func (s RedisStore) Verify(id, answer string, clear bool) bool {
	v := s.Get(id, clear)
	return v == answer
}
