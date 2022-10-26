package delivery

import (
	"log"
	"net/http"
	"project/e-commerce/features/user"
	"project/e-commerce/middlewares"
	"project/e-commerce/utils/helper"
	"strconv"

	// "project/e-commerce/utils/helper"

	"github.com/labstack/echo/v4"
)

type UserDelivery struct {
	userUsecase user.UsecaseInterface
}

func New(e *echo.Echo, usecase user.UsecaseInterface) {
	handler := &UserDelivery{
		userUsecase: usecase,
	}
	e.POST("/register", handler.PostData) // Daftar akun
	e.PUT("/update", handler.UpdateUser, middlewares.JWTMiddleware())
	e.GET("/profile", handler.GetByTokenJWT, middlewares.JWTMiddleware())
	e.GET("/profile/:id", handler.GetByIdWithJWT, middlewares.JWTMiddleware())
	e.DELETE("/profile", handler.DeleteMyAccount, middlewares.JWTMiddleware())

}

func (delivery *UserDelivery) PostData(c echo.Context) error {
	var dataRequest UserRequest
	errBind := c.Bind(&dataRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed To Bind",
		})
	}

	file, _ := c.FormFile("images")
	if file != nil {
		res, err := helper.UploadPosts(c)
		if err != nil {
			return err
		}
		log.Print(res)
		dataRequest.Images = res
	}
	// if errFile != nil {
	// 	return errFile
	// }

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

func (delivery *UserDelivery) UpdateUser(c echo.Context) error {

	var dataUpdate UserRequest
	errBind := c.Bind(&dataUpdate)
	if errBind != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "error Bind data",
		})
	}

	file, _ := c.FormFile("images")
	if file != nil {
		res, err := helper.UploadPosts(c)
		if err != nil {
			return err
		}
		log.Print(res)
		dataUpdate.Images = res
	}
	// if errFile != nil {
	// 	return errFile
	// }

	var add user.Core
	if dataUpdate.Email != "" {
		add.Email = dataUpdate.Email
	}
	if dataUpdate.Name != "" {
		add.Name = dataUpdate.Name
	}
	if dataUpdate.Password != "" {
		add.Password = dataUpdate.Password
	}
	if dataUpdate.Phone != 0 {
		add.Phone = dataUpdate.Phone
	}
	if dataUpdate.Bio != "" {
		add.Bio = dataUpdate.Bio
	}
	if dataUpdate.Saldo != 0 {
		add.Saldo = dataUpdate.Saldo
	}
	if dataUpdate.Images != "" {
		add.Images = dataUpdate.Images
	}

	idToken := middlewares.ExtractToken(c)
	add.ID = uint(idToken)

	row, err := delivery.userUsecase.PutDataId(add)
	if err != nil || row < 1 {
		return c.JSON(400, map[string]interface{}{
			"message": "failed not found",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"message": "success update",
	})

}

func (delivery *UserDelivery) GetByTokenJWT(c echo.Context) error {
	idToken := middlewares.ExtractToken(c)

	res, err := delivery.userUsecase.GetProfile(idToken)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "failed get profile",
		})
	}

	respon := FromCore(res)
	// data := FromUser()
	return c.JSON(200, map[string]interface{}{
		"message": "success get my profile",
		"data":    respon,
		// "product": res,
	})
}

func (delivery *UserDelivery) GetByIdWithJWT(e echo.Context) error {

	idToken := middlewares.ExtractToken(e)
	id, _ := strconv.Atoi(e.Param("id"))
	// id := helper.ParamInt(e)
	if id == -1 {
		return e.JSON(400, map[string]interface{}{
			"message": "failed id not found",
		})
	}

	res, err := delivery.userUsecase.GetDataId(id, idToken)
	if err != nil {
		return e.JSON(400, map[string]interface{}{
			"message": "failed get profile",
		})
	}

	respon := FromCore(res)

	return e.JSON(200, map[string]interface{}{
		"message": "success",
		"data":    respon,
	})

}
func (delivery *UserDelivery) DeleteMyAccount(c echo.Context) error {
	idToken := middlewares.ExtractToken(c)
	row, err := delivery.userUsecase.DeleteAkun(idToken)
	if err != nil || row != 1 {
		return c.JSON(500, map[string]interface{}{
			"message": "failed wrong token",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"message": "success delete account",
	})
}
