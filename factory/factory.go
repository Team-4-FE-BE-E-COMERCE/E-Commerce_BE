package factory

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	authDelivery "project/e-commerce/features/login/delivery"
	loginData "project/e-commerce/features/login/repository"
	authUsecase "project/e-commerce/features/login/service"

	userDelivery "project/e-commerce/features/user/delivery"
	userData "project/e-commerce/features/user/repository"
	userUsecase "project/e-commerce/features/user/service"

	productDelivery "project/e-commerce/features/product/delivery"
	productData "project/e-commerce/features/product/repository"
	productUsecase "project/e-commerce/features/product/service"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {

	userDataFactory := userData.New(db)
	userUsecaseFactory := userUsecase.New(userDataFactory)
	userDelivery.New(e, userUsecaseFactory)

	authDataFactory := loginData.New(db)
	authUsecaseFactory := authUsecase.New(authDataFactory)
	authDelivery.New(e, authUsecaseFactory)

	productDataFactory := productData.New(db)
	productUsecaseFactory := productUsecase.New(productDataFactory)
	productDelivery.New(e, productUsecaseFactory)

}
