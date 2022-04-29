package database

import (
	"answers/pkg/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	//dsn := "root:6379091170@jsk@tcp(49.234.11.169:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", "root:6379091170@jsk@tcp(49.234.11.169:3306)/questions?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Status:", err)
	}

	DB.Debug()
	DB.LogMode(true)
	DB.AutoMigrate(&models.User{}, &models.Question{}, &models.Tag{}, &models.Answer{})
}
