package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-gorm/config"
	"go-gin-gorm/handler"
	"log"
	"net/http"
)

func main() {

	cfg := config.InitConfig()

	r := setRouter()

	log.Fatalln(r.Run(fmt.Sprintf(":%d", cfg.ServerCfg.Port)))
}

func setRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	api := r.Group("api")
	api.GET("/post", handler.GetPost)
	api.POST("/post", handler.SavePost)

	return r
}
