package global

import (
	"fmt"
	"gomod/pkg/setting"
	"time"
)

//服务器配置
type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
}

//数据库配置
type DatabaseSettingS struct {
	DBType   string
	Username string
	Password string
	Host     string
	DBName   string
	Charset  string
}

//定义全局变量
var (
	ServerSetting   *ServerSettingS
	DatabaseSetting *DatabaseSettingS
)

func SetupSetting() error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = s.ReadSection("Database", &DatabaseSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("Server", &ServerSetting)
	if err != nil {
		return err
	}

	fmt.Println("setting:")
	fmt.Println(ServerSetting.HttpPort)
	fmt.Println(DatabaseSetting.DBName)
	return nil
}
