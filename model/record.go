package model

// Record Record
type Record struct {
	ID               int    `json:"id"`
	UserID           string `json:"user_id"`
	ParentID         int    `json:"parent_id,omitempty"`
	TaskID           int    `json:"task_id"`
	SpendTime        int    `json:"spend_time"`
	Date             string `json:"date"`
	CreatedTimestamp int64  `json:"created_timestamp"`
}

// RecordsTable RecordsTable
const RecordsTable = "records"

// RecordsModel RecordsModel
var RecordsModel = RecordsModelObj{
	Table: RecordsTable,
}

// RecordsModelObj RecordsModelObj
type RecordsModelObj struct {
	Table string
}

// Create Create
func (r *RecordsModelObj) Create(row Record) error {
	return DBConn.Table(r.Table).Create(&row).Error
}
