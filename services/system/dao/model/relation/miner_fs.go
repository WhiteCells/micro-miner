package relation

type MinerFs struct {
	MinerID uint `json:"miner_id" gorm:"primaryKey"`
	FsID    uint `json:"fs_id" gorm:"primaryKey"`
}

func (MinerFs) TableName() string {
	return "miner_fs"
}
