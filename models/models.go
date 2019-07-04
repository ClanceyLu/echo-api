package models

import (
	"fmt"
	"time"

	"github.com/ClanceyLu/echo-api/conf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// Model definition
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

var (
	dbConf  *viper.Viper
	appConf *viper.Viper
	db      *gorm.DB
)

func init() {
	dbConf = conf.Conf.Sub("db")
	appConf = conf.Conf.Sub("app")
	if err := connectDB(); err != nil {
		panic(fmt.Errorf("Fatal error connect to db: %s", err))
	}
}

// connectDB 连接数据库
func connectDB() error {
	var (
		err    error
		dbType string = dbConf.GetString("type")
		host   string = dbConf.GetString("host")
		name   string = dbConf.GetString("name")
		user   string = dbConf.GetString("user")
		pwd    string = dbConf.GetString("password")
	)
	if db, err = gorm.Open(dbType,
		fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			user, pwd, host, name)); err != nil {
		return err
	}
	db.LogMode(appConf.GetBool("debug"))
	db.AutoMigrate(&User{})
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	return nil
}
