package global

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DBLink *gorm.DB
)

func SetupDBLink() error {
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=%v&parseTime=True&loc=Local",
		DatabaseSetting.Username,
		DatabaseSetting.Password,
		DatabaseSetting.Host,
		DatabaseSetting.DBName,
		DatabaseSetting.Charset,
	)
	DBLink, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	if ServerSetting.RunMode == "debug" {
		DBLink.Logger.LogMode(logger.LogLevel(2))
	}

	return nil
}
