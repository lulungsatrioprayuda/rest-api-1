package connection

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabseConnection() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/fiber_rest_1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("tidak bisa konek ke database karena" + err.Error())
	}

	fmt.Println("Berhasil konek ke database")
}
