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

func QueryUserWightCon(con string) int {
	sql := fmt.Sprintf("SELECT id FROM users %s;", con)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id
}
