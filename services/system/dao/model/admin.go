package model

type Admin struct {
	PointReward    int `json:"point_reward" gorm:"column:points_reward;type:int"`
	InviteReward   int `json:"invite_reward" gorm:"column:invite_reward;type:int"`
	RegisterSwtich int `json:"register_swtich" gorm:"column:register_swtich;type:int"`
}

func (Admin) TableName() string {
	return "admin"
}
