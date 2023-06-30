package migrations

import (
	"fmt"

	"github.com/lulungsatrioprayuda/rest-api-1/database/connection"
	"github.com/lulungsatrioprayuda/rest-api-1/models"
)

func Migrate() {
	err := connection.DB.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println("Error Migrate")
	}
	fmt.Println("Success Migrate")
}
