package controller

import (
	"CloudRestaurant/tool"
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"

	param_module "CloudRestaurant/param"
)

type CaptchaController struct {
}

func (c CaptchaController) Router(router *gin.RouterGroup) {

	//  http:localhost:8000/memeber/send?phone=12358698745
	router.POST("/getCaptcha", c.generateCaptchaHandler)

	router.POST("/verifyCaptcha", c.captchaVerifyHandle)
}

// base64Captcha create http handler

func (cp CaptchaController) generateCaptchaHandler(context *gin.Context) {
	var store tool.RedisStore = tool.RedisStore{Client: tool.Redis}

	//parse request parameters
	decoder := json.NewDecoder(context.Request.Body)
	var param param_module.CaptchaConfigJsonBody
	err := decoder.Decode(&param)
	if err != nil {
		log.Println(err)
	}
	var driver base64Captcha.Driver

	//choose driver
	switch param.CaptchaType {
	case "audio":
		driver = param.DriverAudio
	case "string":
		driver = param.DriverString.ConvertFonts()
	case "math":
		driver = param.DriverMath.ConvertFonts()
	case "chinese":
		driver = param.DriverChinese.ConvertFonts()
	default:
		driver = param.DriverDigit
	}
	c := base64Captcha.NewCaptcha(driver, &store)
	id, b64s, err := c.Generate()
	body := map[string]interface{}{"code": 1, "data": b64s, "captchaId": id, "msg": "success"}
	if err != nil {
		body = map[string]interface{}{"code": 0, "msg": err.Error()}
	}
	context.JSON(200, body)
}

// base64Captcha verify http handler
func (cp CaptchaController) captchaVerifyHandle(context *gin.Context) {

	var store tool.RedisStore = tool.RedisStore{Client: tool.Redis}

	//parse request parameters
	decoder := json.NewDecoder(context.Request.Body)
	var param param_module.CaptchaCode
	err := decoder.Decode(&param)
	if err != nil {
		log.Println(err)
	}

	//verify the captcha
	body := map[string]interface{}{"code": 0, "msg": "failed"}
	if store.Verify(param.Id, param.VerifyValue, true) {
		body = map[string]interface{}{"code": 1, "msg": "ok"}
	}

	//set json response
	context.JSON(200, body)
}
