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

func (md *MemberDao) InsertCode(sms_coce model.SmsCode) int64 {

	// result :操作行数
	result, err := md.InsertOne(&sms_coce)
	if err != nil {
		log.Println(err.Error())
	}
	return result
}
