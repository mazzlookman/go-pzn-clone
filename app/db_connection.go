package app

import (
	"go-pzn-clone/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnection() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/go_pzn_clone?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)

	return db
}
