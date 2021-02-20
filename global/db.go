package global

import (
	"fmt"
	"gomod/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Db struct {
	Conn *gorm.DB
}

func (d *Db) Connect() error {
	var (
		dbName, user, pwd, host, charset string
	)

	conf := setting.Config.Database
	dbName = conf.Name
	user = conf.User
	pwd = conf.Password
	host = conf.Host
	charset = conf.CharSet

	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=%v&parseTime=True&loc=Local",
		user,
		pwd,
		host,
		dbName,
		charset,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
		return err
	}

	d.Conn = db

	log.Println("数据库连接成功")

	return nil
}

func (d *Db) DB() *gorm.DB {
	_ = d.Connect()
	return d.Conn
}
