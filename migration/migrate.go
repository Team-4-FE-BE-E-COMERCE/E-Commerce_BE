package migration

import (
	cartModel "project/e-commerce/features/cart/repository"
	transactionModel "project/e-commerce/features/order/repository"
	productModel "project/e-commerce/features/product/repository"
	userModel "project/e-commerce/features/user/repository"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&userModel.User{})
	db.AutoMigrate(&productModel.Products{})
	db.AutoMigrate(&cartModel.Cart{})
	db.AutoMigrate(&transactionModel.Transaction{})
	db.AutoMigrate(&transactionModel.Address{})
	db.AutoMigrate(&transactionModel.Payment{})
}
