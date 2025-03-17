package model

type FarmMiner struct {
	FarmID  uint   `json:"farm_id" gorm:"primaryKey"`
	MinerID uint   `json:"miner_id" gorm:"primaryKey"`
	Perm    string `json:"perm" gorm:"column:perm;type:varchar(255)"`
}

func (FarmMiner) TableName() string {
	return "farm_miner"
}
