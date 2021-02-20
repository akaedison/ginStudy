package model

type Movie struct {
	MovieId     uint64 `gorm:"column:movieId;primaryKey;autoIncrement",json:"movieId"`
	MovieName   string `gorm:"column:movieName;size:255",json:"movieName"`
	Playbill    string `gorm:"column:playbill;size:255",json:"playbill"`
	Year        string `gorm:"column:year;size:8",json:"year"`
	Director    string `gorm:"column:director;size:255",json:"director"`
	Writers     string `gorm:"column:writers;size:255",json:"writers"`
	Types       string `gorm:"column:types;size:255",json:"types"`
	Regions     string `gorm:"column:regions;size:255",json:"regions"`
	Languages   string `gorm:"column:languages;size:255",json:"languages"`
	ReleaseInfo string `gorm:"column:releaseInfo;size:255",json:"releaseInfo"`
	Length      string `gorm:"column:length;size:255",json:"length"`
	IMDbLink    string `gorm:"column:IMDbLink;size:255",json:"IMDbLink"`
	AnotherName string `gorm:"column:anotherName;size:255",json:"anotherName"`
}

const tableName = "movie"

func (Movie) TableName() string {
	return tableName
}

/*func CreateTableMovie()  {
	err1 := global.SetupSetting()
	if err1 != nil {
		log.Fatalf("init.setupSetting err: %v", err1)
	}
	err := global.SetupDBLink().AutoMigrate(&Movie{})

	if err != nil {
		panic("创建表失败")
	} else {
		fmt.Printf("创建表%v成功", tableName)
	}
}*/
