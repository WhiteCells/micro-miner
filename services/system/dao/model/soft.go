package model

type Soft struct {
	ID uint `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement"`
}

func (Soft) TableName() string {
	return "soft"
}
