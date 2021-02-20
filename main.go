package main

import (
	"github.com/gin-gonic/gin"
	"gomod/controller"
	"net/http"
)

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

/*func ConnectMySql() *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库初始化失败")
	}
	return db
}*/

func main() {
	//model.CreateTableMovie()
	movie := controller.Movie{}
	movie.CreateTableMovie()
}

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
