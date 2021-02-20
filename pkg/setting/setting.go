package setting

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var (
	Config *Conf
)

/*type Server struct {
	Port			string 			`yaml:"port"`
	ReadTimeout		time.Duration	`yaml:"read-timeout"`
	WriteTimeout	time.Duration	`yaml:"write-timeout"`
}*/

type Database struct {
	Type     string `yaml:"type"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	CharSet  string `yaml:"charset"`
}

type Conf struct {
	/*Server Server `yaml:"server"`*/
	Database Database `yaml:"database"`
}

func init() {
	Config = getConf()
	fmt.Println("配置初始化成功")
}

func getConf() *Conf {
	var c *Conf
	filePath := "./config/config.yaml"
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取文件错误")
	}
	err = yaml.Unmarshal(file, &c)
	if err != nil {
		fmt.Println("读取配置错误")
	}
	return c
}
