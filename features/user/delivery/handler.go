package delivery

import (
	"net/http"
	"project/e-commerce/features/user"

	// "project/e-commerce/utils/helper"

	"github.com/labstack/echo/v4"
)

type UserDelivery struct {
	userUsecase user.UsecaseInterface
}

// ini komen
func New(e *echo.Echo, usecase user.UsecaseInterface) {
	handler := &UserDelivery{
		userUsecase: usecase,
	}
	e.POST("/register", handler.PostData) // Daftar akun

}

func (delivery *UserDelivery) PostData(c echo.Context) error {
	var dataRequest UserRequest
	errBind := c.Bind(&dataRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed To Bind",
		})
	}

	row, err := delivery.userUsecase.PostData(toCore(dataRequest))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed Register",
		})
	}
	if row != 1 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed Register New User",
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Success Registers",
	})
}
