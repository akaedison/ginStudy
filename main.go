package main

import (
	. "fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"time"
)

var (
	db            *gorm.DB = ConnectMySql()
	err           error
	user          string = "root"
	pass          string = "zm159539"
	dbname        string = "go"
	host          string = "47.96.24.50"
	port          string = "3306"
	bookTableName string = "book"
)

type Book struct {
	gorm.Model
	ID     int       `gorm:"primaryKey;autoIncrement";json:"id"`
	Name   string    `gorm:"size:255";json:"name"`
	Price  float64   `json:"price"`
	Author string    `gorm:"size:255";json:"author"`
	Date   time.Time `json:"date"`
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	v1 := r.Group("v1")
	{
		v1.GET("/ping", getV1Ping)
		v1.GET("/hello", getV1Hello)
	}
	v2 := r.Group("v2")
	{
		v2.GET("/ping", getV2Ping)
		v2.GET("/hello", getV2Hello)
	}
	return r
}

func getV1Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func getV1Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "world",
	})
}

func getV2Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "LiPaLa",
	})
}

func getV2Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Edison",
	})
}

func ConnectMySql() *gorm.DB {
	dsn := Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库初始化失败")
	}
	return db
}

func CreateTable() {
	err := db.AutoMigrate(Book{})
	if err != nil {
		panic("创建表失败")
	} else {
		Printf("表%v创建成功", bookTableName)
	}
}

func (b Book) TableName() string {
	return bookTableName
}

func main() {
	//CreateTable()
	//r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	/*err := r.Run(":8080")
	panic(err)*/
	//r := gin.Default()
	/*r.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	go func() {
		//服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}

	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("<Shutdown> Serve ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server ShutDown:", err)
	}

	log.Println("Server exiting")*/
}
