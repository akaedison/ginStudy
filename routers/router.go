package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime/debug"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.NoRoute(HandleNotFound)
	router.NoMethod(HandleNotFound)
	router.Use(Recover)
	return nil
}

func HandleNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "404 not found",
	})
}

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "500 InternalServerError",
			})
		}
	}()

	c.Next()
}
