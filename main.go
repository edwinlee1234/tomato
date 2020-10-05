package main

import (
	"tomato/config"
	"tomato/handler"
	"tomato/middleware"
	"tomato/model"
	"tomato/pkg/redis"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	load()

	gin.SetMode(config.Val.Mode)
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./frontend/dist/", false)))
	r.GET("/ping", handler.Ping)

	api := r.Group("/api")
	{
		api.GET("ouath/google/url", handler.GoogleAccsess)
		api.GET("ouath/google/login", handler.GoogleLogin)
		api.GET("user/info", handler.GetUserInfo)
		api.GET("user/logout", handler.GetUserLogout)
		api.POST("task", middleware.Auth(), handler.CreateTask)
		api.GET("groups", middleware.Auth(), handler.GetGroups)
		api.GET("groups_name", middleware.Auth(), handler.GetGroupsName)
		api.PUT("task", middleware.Auth(), handler.UpdateTask)
		api.PUT("task/done", middleware.Auth(), handler.TaskDone)
		api.DELETE("task/:task_id", middleware.Auth(), handler.DeleteTask)
		api.POST("pomo", middleware.Auth(), handler.Pomo)
		api.GET("report/pie", middleware.Auth(), handler.GetReportPie)
		api.GET("report/line", middleware.Auth(), handler.GetReportLine)
		api.GET("report/task_heatmap", middleware.Auth(), handler.GetTaskHeatMap)
		api.GET("report/group_heatmap", middleware.Auth(), handler.GetAllGroupHeatMap)
		api.GET("day/record", middleware.Auth(), handler.GetDayRecords)
	}

	r.Run(":" + config.Val.Port)

	log.Infof("serve port: %v \n", config.Val.Port)
}

func load() {
	config.Init()
	model.Init()
	redis.Init()
}
