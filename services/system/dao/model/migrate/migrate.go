package migrate

import (
	"micro-system/dao/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		model.Admin{},
		model.Coin{},
		model.Farm{},
		model.Fs{},
		model.Loginlog{},
		model.Miner{},
		model.Operlog{},
		model.Perm{},
		model.Pointlog{},
		model.Pool{},
		model.Role{},
		model.User{},
		model.Wallet{},
	)
}
