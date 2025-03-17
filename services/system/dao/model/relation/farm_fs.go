package relation

type FarmFs struct {
	FarmID uint `json:"farm_id" gorm:"primaryKey"`
	FsID   uint `json:"fs_id" gorm:"primaryKey"`
	Status bool `json:"status" gorm:"column:status;type:tinyint"`
}

func (FarmFs) TableName() string {
	return "farm_fs"
}
