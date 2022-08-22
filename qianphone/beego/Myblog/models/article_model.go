package models

import (
	"Myblog/utils"
)

type Article struct {
	Id         int
	Title      string
	Tags       string
	Short      string
	Content    string
	Author     string
	Createtime int64
	//Status int //Status=0为正常，1为删除，2为冻结
}

// 因为go中没有类class的概念，所以类方法就是不存在的。
// 所以，新增一条数据不是给Article创建一个方法，而是用一个函数，把结构体实例传过去。
// 只有对象方法
func AddArticle(article *Article) (int64, error) {
	sql := `insert into article(
		title,
		tags,
		short,
		content,
		author,
		createtime
		) values (
			?,?,?,?,?,?
		);`
	count, err := utils.ModifyDB(
		sql,
		article.Title,
		article.Tags,
		article.Short,
		article.Content,
		article.Author,
		article.Createtime)
	return count, err
}
