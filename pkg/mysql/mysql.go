package mysql

import (
	"fmt"

	"github.com/ClanceyLu/echo-api/conf"
	"github.com/jinzhu/gorm"

	// gorm 依赖
	_ "github.com/go-sql-driver/mysql"
)

// Client mysql 实例
type Client = *gorm.DB

// Connect 连接数据库，返回 gorm.DB
func Connect() Client {
	var (
		appConf   = conf.Conf.Sub("app")
		mySQLConf = conf.Conf.Sub("mysql")
		host      = mySQLConf.GetString("host")
		user      = mySQLConf.GetString("user")
		password  = mySQLConf.GetString("password")
		name      = mySQLConf.GetString("name")
	)
	db, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			user, password, host, name))
	// defer closeDB(db)
	if err != nil {
		panic(err)
	}
	db.LogMode(appConf.GetBool("debug"))
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	return db
}
