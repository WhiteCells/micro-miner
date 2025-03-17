package relation

type FsPool struct {
	FsID   uint `json:"fs_id" gorm:"primaryKey"`
	PoolID uint `json:"pool_id" gorm:"primaryKey"`
}

func (FsPool) TableName() string {
	return "fs_pool"
}
