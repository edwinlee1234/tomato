package model

import "gorm.io/gorm"

const (
	// TaskDisable TaskDisable
	TaskDisable = 0
	// TaskAble TaskAble
	TaskAble = 1
	// TaskDone TaskDone
	TaskDone = 2
)

// Task Task
type Task struct {
	ID               int    `json:"id,omitempty"`
	ParentID         int    `json:"parent_id,omitempty"`
	UserID           string `json:"user_id,omitempty"`
	Name             string `json:"name,omitempty"`
	Status           int    `json:"status,omitempty"`
	CreatedTimestamp int64  `json:"created_timestamp"`
}

// TasksTable TasksTable
const TasksTable = "tasks"

// TaskModel TaskModel
var TaskModel = TasksModelObj{
	Table: TasksTable,
}

// TasksModelObj TasksModelObj
type TasksModelObj struct {
	Table string
}

// Create Create
func (m TasksModelObj) Create(task Task) error {
	return DBConn.Table(m.Table).Create(&task).Error
}

// GetTasks GetTasks
func (m TasksModelObj) GetTasks(userID string, status []int) ([]*Task, error) {
	var tasks []*Task
	res := DBConn.Table(m.Table).
		Where("user_id = ?", userID).
		Where("status IN (?)", status).
		Order("created_timestamp ASC").
		Scan(&tasks)

	if res.Error == gorm.ErrRecordNotFound {
		return tasks, nil
	}

	return tasks, res.Error
}

// GetGroup GetGroup
func (m TasksModelObj) GetGroup(userID string) ([]*Task, error) {
	var group []*Task
	res := DBConn.Table(m.Table).
		Where("user_id = ?", userID).
		Where("parent_id = ?", 0).
		Where("status = ?", TaskAble).
		Order("created_timestamp ASC").
		Scan(&group)

	if res.Error == gorm.ErrRecordNotFound {
		return group, nil
	}

	return group, res.Error
}
