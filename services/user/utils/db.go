package utils

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Config.MySQL.User,
		Config.MySQL.Password,
		Config.MySQL.Host,
		Config.MySQL.Port,
		Config.MySQL.DBName,
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln("failed to connect db")
		return err
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalln("failed to get database instance")
		return err
	}

	sqlDB.SetMaxIdleConns(Config.MySQL.MaxIdleConns)
	sqlDB.SetMaxOpenConns(Config.MySQL.MaxOpenConns)

	return nil
}
