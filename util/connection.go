package util

import (
	"github.com/sunn789/authInGo/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() error {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	db.AutoMigrate(&user.User{})
	hashedPassword, err := Hasher("test")
	if err != nil {
		return err
	}
	db.Create(&user.User{UserName: "admin", Password: hashedPassword})
	//defer db.cl

	return nil
}
