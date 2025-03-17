package dao

type Loginlog struct {
}

func (Loginlog) TableName() string {
	return "loginlog"
}
