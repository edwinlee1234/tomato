package main

import (
	"tomato/config"
	"tomato/handler"
	"tomato/middleware"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	load()

	gin.SetMode(config.Val.Mode)
	r := gin.Default()
	r.Use(middleware.CROSS())

	r.GET("/ping", handler.Ping)

	r.Run(":" + config.Val.Port)

	log.Infof("serve port: %v \n", config.Val.Port)
}

func load() {
	config.Init()
}
