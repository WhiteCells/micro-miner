package model

type Role struct {
	ID   int    `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"column:name;type:varchar(255);uniqueIndex"`

	// belong User
	Users []User `json:"users" gorm:"many2many:user_role;"`
}

func (Role) TableName() string {
	return "role"
}
