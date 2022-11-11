package service

import (
	"CloudRestaurant/dao"
	"CloudRestaurant/model"
	"CloudRestaurant/param"
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
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
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
	}
	fmt.Printf("response is %#v\n", response)

	// 保存至数据库
	sms_cocde := model.SmsCode{
		Phone:      phone,
		Code:       code,
		BizId:      response.BizId,
		CreateTime: time.Now().Unix(),
	}

	member_dao := dao.MemberDao{tool.DbEngine}
	result := member_dao.InsertCode(sms_cocde)
	return result > 0
}

// 手机验证码登录
func (m MememberService) SmsCodeLogin(p param.SmsLoginParam) *model.Member {
	phone := p.Phone
	code := p.Code

	member_dao := dao.MemberDao{tool.DbEngine}

	// 校验手机和验证
	sms := member_dao.ValidateSmsCode(phone, code)
	if sms.Id == 0 {
		return nil
	}

	// 用户是否存在
	member := member_dao.QueryByPhone(sms.Phone)
	if member.Id != 0 {
		return member
	}

	// 不存在则创建用户
	user := model.Member{}
	user.UserName = phone
	user.Mobile = phone
	user.RegisterTime = time.Now().Unix()

	user.Id = member_dao.InsertMember(user)

	return &user
}
