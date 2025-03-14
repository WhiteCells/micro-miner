package migrate

import (
	"micro-user/dao/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		model.User{},
		model.Role{},
		model.Perm{},
	)
}
