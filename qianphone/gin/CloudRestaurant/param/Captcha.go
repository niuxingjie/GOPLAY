package param

import (
	"github.com/mojocn/base64Captcha"
)

//configJsonBody json request body.
type CaptchaConfigJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

type CaptchaCode struct {
	Id          string `json:"id"`
	VerifyValue string `json:"value"`
}
