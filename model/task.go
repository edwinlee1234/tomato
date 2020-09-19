package model

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
