package handler

import (
	"context"
	"encoding/json"
	"time"
	"tomato/model"
	"tomato/pkg/redis"
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

// GroupsList GroupsList
type GroupsList struct {
	ID    int           `json:"id"`
	Name  string        `json:"name"`
	Tasks []*model.Task `json:"tasks"`
}

// GetGroups GetGroups
func GetGroups(c *gin.Context) {
	userID := c.GetString("user_id")

	var groupList []*GroupsList
	groups, tasks, err := findGroupsAndTasks(userID)
	if err != nil {
		log.WithFields(log.Fields{
			"user_id":    userID,
			"origin_err": err.Error(),
		}).Error("findGroup error")
		res.SystemError(c, res.ErrSystemCode, gin.H{})
		return
	}

	for _, p := range groups {
		groupList = append(groupList, &GroupsList{
			ID:   p.ID,
			Name: p.Name,
		})
	}

	for _, t := range tasks {
		for _, g := range groupList {
			if t.ParentID == g.ID {
				g.Tasks = append(g.Tasks, t)
				break
			}
		}
	}

	res.Success(c, gin.H{
		"groups": groupList,
	})
}
func findGroupsAndTasks(userID string) (groups []*model.Task, tasks []*model.Task, err error) {
	groups, err = findGroups(userID)
	if err != nil {
		return
	}

	parentID := []int{}
	for _, p := range groups {
		parentID = append(parentID, p.ID)
	}

	tasks, err = findTasks(userID, parentID)
	if err != nil {
		return
	}

	return
}

func findGroups(userID string) (groups []*model.Task, err error) {
	val, err := redis.Conn.Get(context.Background(), redisGroupsKey(userID)).Bytes()
	if err == nil {
		if err = json.Unmarshal(val, &groups); err == nil {
			return
		}
	}

	groups, err = model.TaskModel.GetGroup(userID)
	if err != nil {
		return
	}

	groupsJSON, _ := json.Marshal(groups)
	redis.Conn.Set(context.Background(), redisGroupsKey(userID), groupsJSON, 24*3*time.Hour)

	return
}

func findTasks(userID string, parentID []int) (tasks []*model.Task, err error) {
	var data []*model.Task
	redisCache := false
	val, err := redis.Conn.Get(context.Background(), redisTasksKey(userID)).Bytes()
	if err == nil {
		if err = json.Unmarshal(val, &data); err == nil {
			redisCache = true
		}
	}

	if len(data) == 0 {
		data, err = model.TaskModel.GetTasks(userID, []int{model.TaskAble})
		if err != nil {
			return
		}
	}

	if !redisCache {
		tasksJSON, _ := json.Marshal(data)
		redis.Conn.Set(context.Background(), redisTasksKey(userID), tasksJSON, 24*3*time.Hour)
	}

	for _, t := range data {
		for _, pid := range parentID {
			if pid == t.ParentID {
				tasks = append(tasks, t)
				break
			}
		}
	}

	return
}

func redisGroupsKey(userID string) string {
	return "groups:" + userID
}

func redisTasksKey(userID string) string {
	return "tasks:" + userID
}
