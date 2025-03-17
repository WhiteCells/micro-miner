package model

// Miner 属于可以属于多个 Farm
// 但位于不同 Farm 的权限不同
type Miner struct {
	ID     int    `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement"`
	Name   string `json:"name" gorm:"column:name;type:varchar(255)"`
	RigID  string `json:"rig_id" gorm:"column:rig_id;type:varchar(255)"`
	Pass   string `json:"pass" gorm:"column:pass;type:varchar(255)"`
	GpuNum int    `json:"gpu_num" gorm:"column:gpu_num;type:int"`

	Farms Farm `json:"farms" gorm:""`
}

func (Miner) TableName() string {
	return "miner"
}

// Config 可以属于多个 Miner
// Wallet 可以属于多个 Miner
// AutoFan 可以属于多个 Miner
