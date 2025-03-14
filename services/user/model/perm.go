package model

type Perm struct {
	ID   int    `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"column:name;type:varchar(255);uniqueIndex"`

	// belong Role
	Roles []Role `json:"roles" gorm:"many2many:role_perm;"`
}

func (Perm) TableName() string {
	return "perm"
}
