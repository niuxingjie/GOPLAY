package dao

import (
	"CloudRestaurant/model"
	"CloudRestaurant/tool"
	"log"
)

// 这个结构体：初始化时，dao.MemberDao{tool.DbEngine}
type MemberDao struct {
	*tool.Orm
}

// 保存验证码
func (md *MemberDao) InsertCode(sms_coce model.SmsCode) int64 {

	// result :操作行数
	result, err := md.InsertOne(&sms_coce)
	if err != nil {
		log.Println(err.Error())
	}
	return result
}

// 验证验证码
func (md *MemberDao) ValidateSmsCode(phone string, code string) *model.SmsCode {
	var sms_code model.SmsCode

	if _, err := md.Where("phone = ? and code = ?", phone, code).Get(&sms_code); err != nil {
		log.Println(err.Error())
		return nil
	}
	return &sms_code
}

// 根据手机号查询用户
func (md *MemberDao) QueryByPhone(phone string) *model.Member {
	var member model.Member
	if _, err := md.Where("mobile = ?", phone).Get(&member); err != nil {
		log.Println(err.Error())
		return nil
	}
	// 查不到的时候，返回空结构体Id=0：CloudRestaurant/model.Member {Id: 0, UserName: "", Mobile: "", Password: "", RegisterTime: 0, Avatar: "", Balance: 0, IsActive: 0, City: ""}
	return &member
}

// 新建用户
func (md *MemberDao) InsertMember(member model.Member) int64 {

	// result :操作行数
	result, err := md.InsertOne(&member)
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return result
}
