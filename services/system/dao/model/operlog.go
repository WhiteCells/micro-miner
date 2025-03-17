package model

import "time"

type Operlog struct {
	ID     int       `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement"`
	UserID int       `json:"user_id" gorm:"column:user_id;type:int"`
	Time   time.Time `json:"time" gorm:"column:time;type:datetime"`
	Action string    `json:"action" gorm:"column:action;type:varchar(255)"`
	Target string    `json:"target" gorm:"column:target;type:varchar(255)"`
	IP     string    `json:"ip" gorm:"column:ip;type:varchar(255)"`
	Agent  string    `json:"agent" gorm:"column:agent;type:varchar(255)"`
	Status int       `json:"status" gorm:"column:status;type:int"`
	Detail string    `json:"detail" gorm:"column:detail;type:varchar(255)"`
}

func (Operlog) TableName() string {
	return "operlog"
}
