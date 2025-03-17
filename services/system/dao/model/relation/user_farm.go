package relation

type UserFarm struct {
	UserID uint   `json:"user_id" gorm:"primaryKey"`
	FarmID uint   `json:"farm_id" gorm:"primaryKey"`
	Perm   string `json:"perm" gorm:"column:perm;type:varchar(255)"`
}

func (UserFarm) TableName() string {
	return "user_farm"
}
