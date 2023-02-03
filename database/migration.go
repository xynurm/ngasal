package database

import (
	"fmt"
	"ngasal/models"
	"ngasal/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.Post{},
		&models.Category{},
		&models.Campaign{},
		&models.User{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
