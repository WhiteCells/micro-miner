package migrate

import (
	"micro-system/dao/model"
	"micro-system/dao/model/relation"

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
		// relation
		relation.FarmFs{},
		relation.FarmMiner{},
		relation.FsPool{},
		relation.FsSoft{},
		relation.MinerFs{},
		relation.RolePerm{},
		relation.UserFarm{},
		relation.UserRole{},
	)
}
