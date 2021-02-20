package global

import "gorm.io/gorm"

type IDb interface {
	//Connect 初始化数据库配置
	Connect() error
	//DB 返回DB
	DB() *gorm.DB
}
