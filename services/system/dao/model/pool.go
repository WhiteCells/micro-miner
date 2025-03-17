package model

type Pool struct {
	ID   uint   `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"column:name;type:varchar(255)"`
}

func (Pool) TableName() string {
	return "pool"
}
