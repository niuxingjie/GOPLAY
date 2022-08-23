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
	return count, err
}

func FindArticleWithPage(page int) ([]Article, error) {
	num, _ := beego.AppConfig.Int("articleListPageNum")
	page--
	fmt.Println("---------->page", page)
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
	fmt.Println("--------", rows)
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

func ConfigHomeFooterPageCode(page int) HomeFooterPageCode {
	pageCode := HomeFooterPageCode{}
	//查询出总的条数
	num := GetArticleRowsNum()
	//从配置文件中读取每页显示的条数
	pageRow, _ := beego.AppConfig.Int("articleListPageNum")
	//计算出总页数
	allPageNum := (num-1)/pageRow + 1

	pageCode.ShowPage = fmt.Sprintf("%d/%d", page, allPageNum)

	//当前页数小于等于1，那么上一页的按钮不能点击
	if page <= 1 {
		pageCode.HasPre = false
	} else {
		pageCode.HasPre = true
	}

	//当前页数大于等于总页数，那么下一页的按钮不能点击
	if page >= allPageNum {
		pageCode.HasNext = false
	} else {
		pageCode.HasNext = true
	}
	pageCode.PreLink = "/?page=" + strconv.Itoa(page-1)
	pageCode.NextLink = "/?page=" + strconv.Itoa(page+1)
	return pageCode
}

//查询文章的总条数
func GetArticleRowsNum() int {
	row := utils.QueryRowDB("select count(id) from article")
	num := 0
	row.Scan(&num)
	return num
}
