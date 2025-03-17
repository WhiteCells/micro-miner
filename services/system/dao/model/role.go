package model

type Role struct {
	ID   uint   `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"column:name;type:varchar(255);uniqueIndex"`
}

func (Role) TableName() string {
	return "role"
}
