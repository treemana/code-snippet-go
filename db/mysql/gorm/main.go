package gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dsn = "root:this_is_password@tcp(localhost:3306)/this_is_schema?charset=utf8mb4&parseTime=True&loc=Local"
)

func Init() {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Create
	db.Create(&ThisIsModel{Data: "first1"})

	// Read
	var tim ThisIsModel
	db.First(&tim, 1)             // find product with id 1
	db.First(&tim, "id = ?", "1") // find product with code l1212

	// Update - update product's price to 2000
	// db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	// db.Delete(&tim)
}
