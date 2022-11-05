package service

import (
	"CloudRestaurant/tool"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	dysmsapi "github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

type MememberService struct {
}

func (m MememberService) SendCode(phone string) bool {

	// 初始化客户端
	app_config := tool.GetAppConfig()
	client, err := dysmsapi.NewClientWithAccessKey(app_config.RegionID, app_config.AppKey, app_config.AppSecret)
	/* use STS Token
	client, err := dysmsapi.NewClientWithStsToken("cn-qingdao", "<your-access-key-id>", "<your-access-key-secret>", "<your-sts-token>")
	*/
	if err != nil {
		log.Println(err.Error())
		return false
	}

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = phone                   //接收短信的手机号码
	request.SignName = app_config.SignName         //短信签名名称
	request.TemplateCode = app_config.TemplateCode //短信模板ID

	// 生成四位随机验证码
	code := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
	par, err := json.Marshal(map[string]interface{}{
		"code": code,
	})
	if err != nil {
		log.Println(err.Error())
		return false
	}
	request.TemplateParam = string(par)

	// 发送请求
	response, err := client.SendSms(request)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	fmt.Printf("response is %#v\n", response)
	return true
}
