package models

import (
	"Myblog/utils"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
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

	SetArticleRowsNum()
	return count, err
}

func FindArticleWithPage(page int) ([]Article, error) {
	num, _ := beego.AppConfig.Int("articleListPageNum")
	page--
	return QueryArticleWithPage(page, num)
}

func QueryArticleWithPage(page, num int) ([]Article, error) {
	// 第2页的数据：从第num * （2 - 1）个后数num个值
	// 第n页的数据：从第num * （n -1)个后数num个值
	sql := fmt.Sprintf("limit %d,%d", page*num, num)
	return QueryArticlesWithCon(sql)
}

func QueryArticlesWithCon(sql string) ([]Article, error) {
	sql = "select id,title,tags,short,content,author,createtime from article " + sql
	rows, err := utils.QueryDB(sql)
	if err != nil {
		return nil, err
	}
	// fmt.Println("-----文章条数---", rows)
	var artList []Article
	// TODO：相当于python 中while循环，不断的执行rows.Next()直至，返回结果为False，停止执行？
	for rows.Next() {
		id := 0
		title := ""
		tags := ""
		short := ""
		content := ""
		author := ""
		var createtime int64
		createtime = 0
		// TODO:Scan什么用法？
		rows.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
		art := Article{id, title, tags, short, content, author, createtime}
		artList = append(artList, art)
	}
	return artList, nil

}

//  用一个全局变量来存储文章条数
var artcileRowsNum = 0

//只有首次获取行数的时候采取统计表里的行数
func GetArticleRowsNum() int {
	if artcileRowsNum == 0 {
		artcileRowsNum = QueryArticleRowNum()
	}
	return artcileRowsNum
}

func SetArticleRowsNum() {
	artcileRowsNum = QueryArticleRowNum()
}

//查询文章的总条数
func QueryArticleRowNum() int {
	row := utils.QueryRowDB("select count(id) from article")
	num := 0
	row.Scan(&num)
	return num
}

func QueryArticleWithId(id int) Article {
	sql := "select id,title,tags,short,content,author,createtime from article where id=" + strconv.Itoa(id)
	row := utils.QueryRowDB(sql)
	title := ""
	tags := ""
	short := ""
	content := ""
	author := ""
	var createtime int64
	createtime = 0
	row.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
	art := Article{id, title, tags, short, content, author, createtime}
	return art
}

func UpdateArticle(article Article) (int64, error) {
	//数据库操作
	return utils.ModifyDB("update article set title=?,tags=?,short=?,content=? where id=?",
		article.Title, article.Tags, article.Short, article.Content, article.Id)
}

func DeleteArticle(artID int) (int64, error) {
	i, err := deleteArticleWithArtId(artID)
	SetArticleRowsNum()
	return i, err
}
func deleteArticleWithArtId(artID int) (int64, error) {
	return utils.ModifyDB("delete from article where id=?", artID)
}
