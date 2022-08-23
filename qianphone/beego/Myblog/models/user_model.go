package models

import (
	"fmt"

	"Myblog/utils"
)

type User struct {
	Id         int
	Username   string
	Password   string
	Status     int // 0正常 1删除
	Createtime int64
}

func InsertUser(user User) (int64, error) {
	sql := `INSERT INTO users(
		username,
		password,
		status,
		createtime
	) values (
		?,
		?,
		?,
		?
	);`
	return utils.ModifyDB(sql, user.Username, user.Password, user.Status, user.Createtime)
}

// 拼接过滤条件查询用户id
func QueryUserWightCon(con string) int {
	sql := fmt.Sprintf("SELECT id FROM users %s;", con)
	// fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id
}

// 根据用户名来查询用户
func QueryUserWithUsername(username string) int {
	sql := fmt.Sprintf("where username='%s'", username)
	return QueryUserWightCon(sql)
}

// 根据用户和密码查询用户
func QueryUserWithParam(username, password string) int {
	sql := fmt.Sprintf("where username='%s' and password='%s'", username, password)
	return QueryUserWightCon(sql)
}
