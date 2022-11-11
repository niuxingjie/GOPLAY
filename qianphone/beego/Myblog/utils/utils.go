package utils

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/astaxie/beego"

	// mysql驱动包
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	InitMysql()
}

func InitMysql() {
	fmt.Println("InitMysql......")
	driverName := beego.AppConfig.String("driverName")

	// orm.RegisterDriver(driverName, orm.DRMySQL)

	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")

	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

	// err := orm.RegisterDataBase("default", driverName, dbConn)
	db1, err := sql.Open(driverName, dbConn)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		db = db1
		CreateTableWithUser()
		CreateTableWithArticle()
		CreateTableWithAlbum()
	}

}

func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		status INT(4),
		createtime INT(10)
	);`
	ModifyDB(sql)
}

//创建文章表
func CreateTableWithArticle() {
	sql := `create table if not exists article(
		id int(4) primary key auto_increment not null,
		title varchar(30),
		author varchar(20),
		tags varchar(30),
		short varchar(255),
		content longtext,
		createtime int(10)
		);`
	ModifyDB(sql)
}

func CreateTableWithAlbum() {
	sql := `create table if not exists album(
		id int(4) primary key auto_increment not null,
		filepath varchar(255),
		filename varchar(64),
		status int(4),
		createtime int(10)
		);`
	ModifyDB(sql)
}

func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}

func MD5(str string) string {
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5str
}

func QueryDB(sql string) (*sql.Rows, error) {
	return db.Query(sql)
}

func SwitchTimeStampToData(timestamp int64) string {
	return strconv.FormatInt(timestamp, 10)
}
