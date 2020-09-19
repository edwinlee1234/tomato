package handler

import (
	"time"
	"tomato/model"
	"tomato/pkg/res"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Ping Ping
func Ping(c *gin.Context) {
	res.Success(c, gin.H{
		"msg": "pong",
	})
}

// CreateTask CreateTask
func CreateTask(c *gin.Context) {
	var taskData model.Task
	c.BindJSON(&taskData)

	taskData.UserID = c.GetString("user_id")
	taskData.Status = model.TaskAble
	taskData.CreatedTimestamp = time.Now().UTC().Unix()

	if err := model.TaskModel.Create(taskData); err != nil {
		log.WithFields(log.Fields{
			"task_data":  taskData,
			"origin_err": err.Error(),
		}).Error("db error")
		res.SystemError(c, res.ErrSystemCode, gin.H{})
		return
	}

	res.Success(c, gin.H{})
}
