package model

type Perm struct {
	ID   uint   `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"column:name;type:varchar(255);uniqueIndex"`
}

func (Perm) TableName() string {
	return "perm"
}
