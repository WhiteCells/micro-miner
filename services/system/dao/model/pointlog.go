package model

import "time"

type Pointlog struct {
	ID      int       `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement"`
	UserID  int       `json:"user_id" gorm:"column:user_id;type:int"`
	Type    string    `json:"type" gorm:"column:type;type:varchar(255)"`
	Amount  float32   `json:"amount" gorm:"column:amount;type:float"`
	Balance float32   `json:"balance" gorm:"column:balance;type:float"`
	Time    time.Time `json:"time" gorm:"column:time;type:datetime"`
	Detail  string    `json:"detail" gorm:"column:detail;type:varchar(255)"`
}

func (Pointlog) TableName() string {
	return "pointlog"
}
