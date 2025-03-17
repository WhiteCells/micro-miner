package model

type Admin struct {
	PointReward    uint    `json:"point_reward" gorm:"column:points_reward;type:int"`
	RechargeRatio  float32 `json:"recharge_ratio" gorm:"column:recharge_ratio;type:float"`
	InviteReward   uint    `json:"invite_reward" gorm:"column:invite_reward;type:int"`
	RegisterSwtich uint    `json:"register_swtich" gorm:"column:register_swtich;type:int"`
}

func (Admin) TableName() string {
	return "admin"
}
