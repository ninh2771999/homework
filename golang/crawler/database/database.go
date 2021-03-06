package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDb() *gorm.DB {
	info := "root:password@tcp(127.0.0.1:3306)/testdb"
	db, err := gorm.Open(mysql.Open(info))

	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to the database!")
	}
	return db
}