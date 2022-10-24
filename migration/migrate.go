package migration

import (
	userModel "project/e-commerce/features/user/repository"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&userModel.User{})

}
