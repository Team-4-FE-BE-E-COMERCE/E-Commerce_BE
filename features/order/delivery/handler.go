package delivery

import (
	"net/http"
	"project/e-commerce/features/order"
	"project/e-commerce/middlewares"

	"github.com/labstack/echo/v4"
)

type TransactionDelivery struct {
	transactionUsecase order.UsecaseInterface
}

func New(e *echo.Echo, usecase order.UsecaseInterface) {
	handler := &TransactionDelivery{
		transactionUsecase: usecase,
	}

	e.POST("/orders", handler.PostDataOrders, middlewares.JWTMiddleware())
	e.PUT("/orders/confirm", handler.PutStatusConfirm, middlewares.JWTMiddleware())
	e.PUT("/orders/cancel", handler.PutDeleteOrder, middlewares.JWTMiddleware())
	e.GET("/orders", handler.GetOrderHistory, middlewares.JWTMiddleware())

}

func (delivery *TransactionDelivery) PostDataOrders(c echo.Context) error {

	idtoken := middlewares.ExtractToken(c)
	var data Request
	errBind := c.Bind(&data)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "error bind data",
		})
	}

	requestAddress, requestPayment := data.fromCore()

	row, err := delivery.transactionUsecase.PostData(idtoken, requestAddress, requestPayment)
	if err != nil || row == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed checkout",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes status waiting",
	})

}

func (delivery *TransactionDelivery) PutDeleteOrder(c echo.Context) error {

	idtoken := middlewares.ExtractToken(c)

	row, err := delivery.transactionUsecase.DeleteOrder(idtoken, "cancel")
	if err != nil || row == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to canceled orders",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success cancel order",
	})

}

func (delivery *TransactionDelivery) GetOrderHistory(c echo.Context) error {

	idToken := middlewares.ExtractToken(c)

	data, err := delivery.transactionUsecase.GetOrder(idToken)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to get history orders",
		})
	}
	res := toRespon(data)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "succes get history order",
		"your history": res,
	})
}
func (delivery *TransactionDelivery) PutStatusConfirm(c echo.Context) error {

	idtoken := middlewares.ExtractToken(c)

	row, err := delivery.transactionUsecase.PutStatus(idtoken, "confirm")
	if err != nil || row == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "low stock product",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success status confirmed",
	})

}
