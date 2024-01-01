package Seeders

import (
	"app/Models"
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

func seedUsers(DB *gorm.DB) {
	DB.Create(&Models.User{
		Username: gofakeit.Name(),
		Email:    gofakeit.Email(),
		Password: gofakeit.Password(true, true, true, true, true, 10),
		Group:    "admin",
		Token:    gofakeit.Email(),
	})
}
