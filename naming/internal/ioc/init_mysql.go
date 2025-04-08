package ioc

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"naming/naming/internal/repo/dao"
)

func InitGorm(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&dao.Naming{})
	if err != nil {
		panic(err)
	}
	return db
}
