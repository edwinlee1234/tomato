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

// GetByUserIDAndTimeBetween GetByUserIDAndTimeBetween
func (r *RecordsModelObj) GetByUserIDAndTimeBetween(userID string, startDate, endDate string) ([]Record, error) {
	var data []Record

	res := DBConn.Table(r.Table).
		Where("user_id = ?", userID).
		Where("date BETWEEN ? AND ?", startDate, endDate).
		Scan(&data)

	return data, res.Error
}

// SumByDate SumByDate
func (r *RecordsModelObj) SumByDate(userID string, startDate, endDate string) ([]Record, error) {
	var data []Record

	res := DBConn.Table(r.Table).
		Select("date, SUM(spend_time) as spend_time").
		Where("user_id = ?", userID).
		Where("date BETWEEN ? AND ?", startDate, endDate).
		Group("date").
		Order("date ASC").
		Scan(&data)

	return data, res.Error
}

// SumByDateByParentID SumByDateByParentID
func (r *RecordsModelObj) SumByDateByParentID(userID string, parentID int, startDate, endDate string) ([]Record, error) {
	var data []Record

	res := DBConn.Table(r.Table).
		Select("date, task_id, SUM(spend_time) as spend_time").
		Where("user_id = ?", userID).
		Where("date BETWEEN ? AND ?", startDate, endDate).
		Group("date, task_id").
		Order("date ASC").
		Scan(&data)

	return data, res.Error
}

// SumByDateGroupByParent SumByDateGroupByParent
func (r *RecordsModelObj) SumByDateGroupByParent(userID string, startDate, endDate string) ([]Record, error) {
	var data []Record

	res := DBConn.Table(r.Table).
		Select("date, parent_id, SUM(spend_time) as spend_time").
		Where("user_id = ?", userID).
		Where("date BETWEEN ? AND ?", startDate, endDate).
		Group("date, parent_id").
		Order("date ASC").
		Scan(&data)

	return data, res.Error
}
