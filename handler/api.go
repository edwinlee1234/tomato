package handler

import (
	"context"
	"encoding/json"
	"strconv"
	"time"
	"tomato/config"
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

	if taskData.ParentID == 0 {
		redis.Conn.Del(context.Background(), redisGroupsKey(c.GetString("user_id")))
	} else {
		redis.Conn.Del(context.Background(), redisTasksKey(c.GetString("user_id")))
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

// UpdateTask UpdateTask
func UpdateTask(c *gin.Context) {
	var params struct {
		ID       int    `json:"id" binding:"required"`
		TaskName string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&params); err != nil {
		res.BadRequest(c, res.ErrParamsCode, gin.H{
			"msg": err.Error(),
		})
		return
	}

	task := model.Task{
		Name: params.TaskName,
	}
	if err := model.TaskModel.Update(params.ID, c.GetString("user_id"), task); err != nil {
		log.WithFields(log.Fields{
			"origin_err": err.Error(),
		}).Error("db error")
		res.SystemError(c, res.ErrSystemCode, gin.H{})
		return
	}

	redis.Conn.Del(context.Background(), redisGroupsKey(c.GetString("user_id")))
	redis.Conn.Del(context.Background(), redisTasksKey(c.GetString("user_id")))

	res.Success(c, gin.H{})
}

// DeleteTask DeleteTask
func DeleteTask(c *gin.Context) {
	taskID := c.Param("task_id")
	if c.Param("task_id") == "" {
		res.BadRequest(c, res.ErrParamsCode, gin.H{
			"msg": "id not found",
		})
		return
	}

	id, err := strconv.Atoi(taskID)
	if err != nil {
		res.BadRequest(c, res.ErrParamsCode, gin.H{
			"msg": "id error",
		})
		return
	}

	if err := model.TaskModel.Delete(id); err != nil {
		log.WithFields(log.Fields{
			"origin_err": err.Error(),
		}).Error("db error")
		res.SystemError(c, res.ErrSystemCode, gin.H{})
		return
	}

	redis.Conn.Del(context.Background(), redisGroupsKey(c.GetString("user_id")))
	redis.Conn.Del(context.Background(), redisTasksKey(c.GetString("user_id")))

	res.Success(c, gin.H{})
}

// Pomo Pomo
func Pomo(c *gin.Context) {
	var params struct {
		ID   int `json:"id" binding:"required"`
		Time int `json:"time" binding:"required"`
	}

	if err := c.ShouldBindJSON(&params); err != nil {
		res.BadRequest(c, res.ErrParamsCode, gin.H{
			"msg": err.Error(),
		})
		return
	}

	userID := c.GetString("user_id")

	// 檢查任務存不存在
	task, err := model.TaskModel.GetByIDAndUserID(params.ID, userID)
	if err != nil || task.ID == 0 {
		res.BadRequest(c, res.ErrParamsCode, gin.H{
			"msg": "has no task",
		})
		return
	}

	nowTime := time.Now()
	record := model.Record{
		UserID:           userID,
		TaskID:           task.ID,
		ParentID:         task.ParentID,
		SpendTime:        params.Time,
		Date:             nowTime.Format(config.Val.TimeFormat),
		CreatedTimestamp: nowTime.UTC().Unix(),
	}
	if err := model.RecordsModel.Create(record); err != nil {
		log.WithFields(log.Fields{
			"origin_err": err.Error(),
		}).Error("db error")
		res.SystemError(c, res.ErrSystemCode, gin.H{})
		return
	}

	res.Success(c, gin.H{})
}

// TaskReport TaskReport
type TaskReport struct {
	ID             int    `json:"id"`
	ParentID       int    `json:"-"`
	Name           string `json:"name"`
	TotalSpendTime int    `json:"total_spend_time"`
}

// GroupReport GroupReport
type GroupReport struct {
	ID             int           `json:"id"`
	Name           string        `json:"name"`
	TotalSpendTime int           `json:"total_spend_time"`
	Tasks          []*TaskReport `json:"tasks"`
}

// GetReportPie GetReportPie
func GetReportPie(c *gin.Context) {
	userID := c.GetString("user_id")

	nowTime := time.Now()
	endDate := nowTime.Format(config.Val.TimeFormat)
	startDate := nowTime.AddDate(0, 0, -7).Format(config.Val.TimeFormat)

	// 取得記錄的時間加總
	records, err := model.RecordsModel.GetByUserIDAndTimeBetween(userID, startDate, endDate)
	if err != nil {
		log.WithFields(log.Fields{
			"origin_err": err.Error(),
		}).Error("GetByUserIDAndTimeBetween error")
		res.SystemError(c, res.ErrSystemCode, gin.H{})
		return
	}

	var tasksReport []*TaskReport
	var groupsReport []*GroupReport

	groups, tasks, err := findGroupsAndTasks(userID)
	if err != nil {
		log.WithFields(log.Fields{
			"user_id":    userID,
			"origin_err": err.Error(),
		}).Error("findGroup error")
		res.SystemError(c, res.ErrSystemCode, gin.H{})
		return
	}

	for _, g := range groups {
		groupsReport = append(groupsReport, &GroupReport{
			ID:   g.ID,
			Name: g.Name,
		})
	}

	for _, t := range tasks {
		tasksReport = append(tasksReport, &TaskReport{
			ID:       t.ID,
			ParentID: t.ParentID,
			Name:     t.Name,
		})
	}

	// 找到記錄對應的任務
	for _, r := range records {
		for _, t := range tasksReport {
			if t.ID == r.TaskID {
				t.TotalSpendTime += r.SpendTime
				break
			}
		}
	}

	// 找到任務對應的群組
	for _, t := range tasksReport {
		for _, g := range groupsReport {
			if t.ParentID == g.ID {
				g.TotalSpendTime += t.TotalSpendTime
				g.Tasks = append(g.Tasks, t)
				break
			}
		}
	}

	res.Success(c, gin.H{
		"report": groupsReport,
	})
}

// DateReport DateReport
type DateReport struct {
	Date           string `json:"date"`
	TotalSpendTime int    `json:"total_spend_time"`
}

func rangeDate(endDate time.Time, days int) []string {
	dates := []string{}
	for i := days - 1; i > 0; i-- {
		dates = append(dates, endDate.AddDate(0, 0, -i).Format(config.Val.TimeFormat))
	}

	dates = append(dates, endDate.Format(config.Val.TimeFormat))

	return dates
}

// GetReportLine GetReportLine
func GetReportLine(c *gin.Context) {
	userID := c.GetString("user_id")

	nowTime := time.Now()
	endDate := nowTime.Format(config.Val.TimeFormat)
	startDate := nowTime.AddDate(0, 0, -7).Format(config.Val.TimeFormat)

	dateStatistics := rangeDate(nowTime, 7)
	daysReport, err := sumDateReport(userID, startDate, endDate, dateStatistics)
	if err != nil {
		log.WithFields(log.Fields{
			"origin_err": err.Error(),
		}).Error("db error")
		res.SystemError(c, res.ErrSystemCode, gin.H{})
		return
	}

	res.Success(c, gin.H{
		"report": daysReport,
	})
}

func sumDateReport(userID, startDate, endDate string, dateStatistics []string) ([]*DateReport, error) {
	var daysReport []*DateReport
	daySum, err := model.RecordsModel.SumByDate(userID, startDate, endDate)
	if err != nil {
		return daysReport, err
	}
	// 初始化7天的資料
	for _, ds := range dateStatistics {
		daysReport = append(daysReport, &DateReport{
			Date: ds,
		})
	}

	for _, d := range daySum {
		// 轉換DB日期的型態
		t1, err := time.Parse(time.RFC3339, d.Date)
		if err != nil {
			return daysReport, err
		}
		parsedDate := t1.Format("2006-01-02")

		// 把值加總在對應的日期上
		for _, dr := range daysReport {
			if dr.Date == parsedDate {
				dr.TotalSpendTime = d.SpendTime
				break
			}
		}
	}

	return daysReport, nil
}

// HeatMapReport HeatMapReport
type HeatMapReport struct {
	ID             int    `json:"-"`
	Name           string `json:"name"`
	Date           string `json:"date"`
	TotalSpendTime int    `json:"total_spend_time"`
}

// GetTaskHeatMap GetTaskHeatMap
func GetTaskHeatMap(c *gin.Context) {
	var heatmap []*HeatMapReport
	userID := c.GetString("user_id")
	groupIDstr, _ := c.GetQuery("group_id")

	groupID, err := strconv.Atoi(groupIDstr)
	if err != nil || groupID == 0 {
		res.BadRequest(c, res.ErrParamsCode, gin.H{"msg": "group_id not found"})
		return
	}

	nowTime := time.Now()
	endDate := nowTime.Format(config.Val.TimeFormat)
	startDate := nowTime.AddDate(0, 0, -7).Format(config.Val.TimeFormat)
	dateStatistics := rangeDate(nowTime, 7)

	tasks, err := model.TaskModel.GetTasksByParentID([]int{groupID}, userID, []int{model.TaskAble})
	if err != nil {
		log.WithFields(log.Fields{
			"user_id":    userID,
			"origin_err": err.Error(),
		}).Error("GetTasksByParentID error")
		res.SystemError(c, res.ErrSystemCode, gin.H{})
		return
	}

	data, err := model.RecordsModel.SumByDateByParentID(userID, groupID, startDate, endDate)
	if err != nil {
		log.WithFields(log.Fields{
			"user_id":    userID,
			"origin_err": err.Error(),
		}).Error("SumByDateByParentID error")
		res.SystemError(c, res.ErrSystemCode, gin.H{})
		return
	}

	for _, t := range tasks {
		for _, ds := range dateStatistics {
			heatmap = append(heatmap, &HeatMapReport{
				ID:   t.ID,
				Name: t.Name,
				Date: ds,
			})
		}
	}

	for _, d := range data {
		t1, err := time.Parse(time.RFC3339, d.Date)
		if err != nil {
			log.WithFields(log.Fields{
				"date":       d.Date,
				"origin_err": err.Error(),
			}).Error("GetAllGroupHeatMap Parse error")
			res.SystemError(c, res.ErrSystemCode, gin.H{})
			return
		}

		parsedDate := t1.Format("2006-01-02")

		for _, h := range heatmap {
			if h.ID == d.TaskID && h.Date == parsedDate {
				h.TotalSpendTime = d.SpendTime
				break
			}
		}
	}

	res.Success(c, gin.H{
		"report": heatmap,
	})
}

// GetAllGroupHeatMap GetAllGroupHeatMap
func GetAllGroupHeatMap(c *gin.Context) {
	var heatmap []*HeatMapReport
	userID := c.GetString("user_id")

	nowTime := time.Now()
	endDate := nowTime.Format(config.Val.TimeFormat)
	startDate := nowTime.AddDate(0, 0, -7).Format(config.Val.TimeFormat)
	dateStatistics := rangeDate(nowTime, 7)

	data, err := model.RecordsModel.SumByDateGroupByParent(userID, startDate, endDate)
	if err != nil {
		log.WithFields(log.Fields{
			"user_id":    userID,
			"origin_err": err.Error(),
		}).Error("SumByDateByParent error")
		res.SystemError(c, res.ErrSystemCode, gin.H{})
		return
	}

	groups, err := model.TaskModel.GetGroup(userID)
	if err != nil {
		log.WithFields(log.Fields{
			"user_id":    userID,
			"origin_err": err.Error(),
		}).Error("GetGroup error")
		res.SystemError(c, res.ErrSystemCode, gin.H{})
		return
	}

	for _, g := range groups {
		for _, ds := range dateStatistics {
			heatmap = append(heatmap, &HeatMapReport{
				ID:   g.ID,
				Name: g.Name,
				Date: ds,
			})
		}
	}

	for _, d := range data {
		t1, err := time.Parse(time.RFC3339, d.Date)
		if err != nil {
			log.WithFields(log.Fields{
				"date":       d.Date,
				"origin_err": err.Error(),
			}).Error("GetAllGroupHeatMap Parse error")
			res.SystemError(c, res.ErrSystemCode, gin.H{})
			return
		}

		parsedDate := t1.Format("2006-01-02")

		for _, h := range heatmap {
			if h.ID == d.ParentID && h.Date == parsedDate {
				h.TotalSpendTime = d.SpendTime
				break
			}
		}
	}

	res.Success(c, gin.H{
		"report": heatmap,
	})
}
