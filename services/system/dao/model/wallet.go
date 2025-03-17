package model

type Wallet struct {
	ID      int    `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement"`
	Name    string `json:"name" gorm:"column:name;type:varchar(255)"`
	Address string `json:"address" gorm:"column:address;type:varchar(255)"`
}

func (Wallet) TableName() string {
	return "wallet"
}
