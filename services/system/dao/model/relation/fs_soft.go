package relation

type FsSoft struct {
	FsID uint `json:"fs_id" gorm:"primaryKey"`
	Soft uint `json:"soft" gorm:"primaryKey"`
}

func (FsSoft) TableName() string {
	return "fs_soft"
}
