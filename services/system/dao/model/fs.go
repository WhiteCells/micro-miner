package model

type Fs struct {
	ID uint `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement"`
}

func (Fs) TableName() string {
	return "fs"
}

// fs 和 coin 关联
// fs 和 pool 关联
// fs 和 soft 关联
