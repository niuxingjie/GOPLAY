package tool

import (
	"CloudRestaurant/model"

	"github.com/go-xorm/xorm"

	_ "github.com/go-sql-driver/mysql"
)

type Orm struct {
	*xorm.Engine
}

var DbEngine *Orm

func OrmEngine(conf *Config) (*Orm, error) {
	database_config := conf.Database

	// 格式："mysql://root:password@tcp(localhost:3306)/texo_ewallet_db?charset=utf8"
	conn := database_config.Username + ":" + database_config.Password + "@tcp(" + database_config.Host + ":" + database_config.Port + ")/" + database_config.DbName + "?charset=" + database_config.Charset

	// xorm的接口：建立数据库连接然后赋值给orm引擎engine属性值
	engine, err := xorm.NewEngine(database_config.Driver, conn)
	if err != nil {
		return nil, err
	}

	// 是否打印sql语句配置
	engine.ShowSQL(database_config.ShowSQL)

	// 表映射同步
	err = engine.Sync2(new(model.SmsCode))
	if err != nil {
		return nil, err
	}

	orm := new(Orm)
	orm.Engine = engine

	DbEngine = orm
	return orm, err
}
