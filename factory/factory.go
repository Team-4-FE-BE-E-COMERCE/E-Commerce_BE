package factory

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	userDelivery "project/e-commerce/features/user/delivery"
	userData "project/e-commerce/features/user/repository"
	userUsecase "project/e-commerce/features/user/service"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {

	userDataFactory := userData.New(db)
	userUsecaseFactory := userUsecase.New(userDataFactory)
	userDelivery.New(e, userUsecaseFactory)

}
