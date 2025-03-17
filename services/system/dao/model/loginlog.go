package model

import "time"

type Loginlog struct {
	ID     int       `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement"`
	UserID int       `json:"user_id" gorm:"column:user_id;type:int"`
	Time   time.Time `json:"time" gorm:"column:time;type:datetime"`
	IP     string    `json:"ip" gorm:"column:ip;type:varchar(255)"`
	Status int       `json:"status" gorm:"column:status;type:int"`
}

func (Loginlog) TableName() string {
	return "loginlog"
}
