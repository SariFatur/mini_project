package config

import (
	"myproject/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:yogakeroh14@tcp(localhost:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	initMigrate()
}

func initMigrate() {
	DB.AutoMigrate(&model.Guest{})
}

func InitDBTest() {
	dsn := "root:yogakeroh14@tcp(localhost:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	initMigrate()
}

func initMigrateTest() {
	DB.Migrator().DropTable(&model.Guest{})
	DB.AutoMigrate(&model.Guest{})
}
