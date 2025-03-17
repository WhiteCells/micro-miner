package model

import "time"

type Farm struct {
	ID       uint      `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement"`
	Name     string    `json:"name" gorm:"column:name;type:varchar(255)"`
	TimeZone time.Time `json:"time_zone" gorm:"column:time_zone;type:datetime"`
	Hash     string    `json:"hash" gorm:"column:hash;type:varchar(255)"`
	MinerNum uint      `json:"miner_num" gorm:"column:miner_num;type:int"`
	GpuNum   uint      `json:"gpu_num" gorm:"column:gpu_num;type:int"`
}

func (Farm) TableName() string {
	return "farm"
}
