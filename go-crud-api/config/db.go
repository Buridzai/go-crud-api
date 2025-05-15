package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=buri114455 dbname=go_crud port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("⚠️ Không thể kết nối DB")
	}
	fmt.Println("✅ Kết nối DB thành công")
	DB = db
}
