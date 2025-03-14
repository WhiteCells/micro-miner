package model

import "time"

type User struct {
	ID             int       `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement"`
	Name           string    `json:"name" gorm:"column:name;type:varchar(255)"`
	Email          string    `json:"email" gorm:"column:email;type:varchar(255)"`
	Password       string    `json:"password" gorm:"column:password;type:varchar(255)"`
	Secret         string    `json:"secret" gorm:"column:secret;type:varchar(255)"`
	InvitePoints   float32   `json:"invite_points" gorm:"column:invite_points;type:float"`
	RechargePoints float32   `json:"recharge_points" gorm:"column:recharge_points;type:float"`
	LastBalance    float32   `json:"last_balance" gorm:"column:last_balance;type:float"`
	Status         int       `json:"status" gorm:"column:status;type:int"`
	LastLoginAt    time.Time `json:"last_login_at" gorm:"column:last_login_at;type:datetime"`
	LastLoginIp    string    `json:"last_login_ip" gorm:"column:last_login_ip;type:varchar(255)"`
	InviteCode     string    `json:"invite_code" gorm:"column:invite_code;type:varchar(255)"`
	InviteBy       int       `json:"invite_by" gorm:"column:invite_by;type:int"`
	OwnerFarmNum   int       `json:"owner_farm_num" gorm:"column:owner_farm_num;type:int"`
	OwnerMinerNum  int       `json:"owner_miner_num" gorm:"column:owner_miner_num;type:int"`
}

func (User) TableName() string {
	return "user"
}
