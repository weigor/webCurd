package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	//DB.AutoMigrate(&model.Order{})
	//DB.SingularTable(true)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}
func Close() {
	DB.Close()
}
