package delivery

import (
	"net/http"
	"project/e-commerce/features/cart"
	"project/e-commerce/middlewares"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CartDelivery struct {
	cartUsecase cart.UsecaseInterface
}

func New(e *echo.Echo, usecase cart.UsecaseInterface) {
	handler := &CartDelivery{
		cartUsecase: usecase,
	}
	e.POST("/carts", handler.PostData, middlewares.JWTMiddleware())
	e.GET("/carts", handler.GetAllCart, middlewares.JWTMiddleware())
	e.PUT("/carts/:id", handler.UpdateCart, middlewares.JWTMiddleware())
	e.DELETE("/carts/:id", handler.DeleteCart, middlewares.JWTMiddleware())
}

func (delivery *CartDelivery) PostData(c echo.Context) error {
	idToken := middlewares.ExtractToken(c)
	if idToken == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Failed Token Not Found",
		})
	}

	var dataCart CartRequest
	errBind := c.Bind(&dataCart)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "error bind data",
		})
	}

	dataCart.Quantity = 1
	dataCart.UserID = idToken

	row, err := delivery.cartUsecase.PostData(toCore(dataCart))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error insert data",
		})
	}
	if row == 2 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "product add to cart again",
		})
	} else if row == 3 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed add to cart",
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "succes add to cart",
	})
}

func (delivery *CartDelivery) DeleteCart(c echo.Context) error {
	idToken := middlewares.ExtractToken(c)

	id := c.Param("id")
	idCnv, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "param must be number",
		})
	}

	row, err := delivery.cartUsecase.DeleteCart(idToken, idCnv)
	if err != nil || row != 1 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Row Not Found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Delete Cart",
	})
}

func (delivery *CartDelivery) GetAllCart(c echo.Context) error {

	idToken := middlewares.ExtractToken(c)
	data, err := delivery.cartUsecase.GetByToken(idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed get cart",
		})
	} else if len(data) == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "you dont have product in cart",
		})
	}

	dataCart, totalRes := fromCoreList(data)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Get Your Cart",
		"cart":    dataCart,
		"total":   totalRes,
	})

}

func (delivery *CartDelivery) UpdateCart(c echo.Context) error {

	id := c.Param("id")
	idCart, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed Error Internal",
		})
	}

	update := c.QueryParam("update")
	if update == "increment" || update == "decrement" {
		cartID := idCart
		idToken := middlewares.ExtractToken(c)

		row, err := delivery.cartUsecase.PutData(cartID, idToken, update)
		if err != nil || row < 1 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Error bad Request",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success Update cart",
		})
	} else {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "query param must be increment or decrement",
		})
	}

}
