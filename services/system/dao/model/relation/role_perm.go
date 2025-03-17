package model

type RolePerm struct {
	RoleID uint `json:"role_id" gorm:"primaryKey"`
	PermID uint `json:"perm_id" gorm:"primaryKey"`
}

func (RolePerm) TableName() string {
	return "role_perm"
}
