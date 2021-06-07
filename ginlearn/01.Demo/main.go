package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		cCp := c.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			cCp .String(http.StatusOK, "Welcome Gin Server")
		}()

	})

	router.Run()
}
