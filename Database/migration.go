package Database

import (
	"app/Models"
	"gorm.io/gorm"
)

func Migrate(DB *gorm.DB) {
	DB.AutoMigrate(
		&Models.User{},
	)
}
