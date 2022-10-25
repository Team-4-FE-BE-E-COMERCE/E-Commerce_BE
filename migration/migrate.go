package migration

import (
	productModel "project/e-commerce/features/product/repository"
	userModel "project/e-commerce/features/user/repository"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&userModel.User{})
	db.AutoMigrate(&productModel.Products{})
}
