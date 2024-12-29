package model

type Todo struct {
	Id          int    `json:"task_id" gorm:"primary key"`
	Data        string `json:"task_name"`
	IsCompleted string `json:"status"`
}
