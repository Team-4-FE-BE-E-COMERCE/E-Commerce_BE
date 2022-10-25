package delivery

import (
	"net/http"
	cf "project/e-commerce/config"
	"project/e-commerce/features/product"
	"project/e-commerce/middlewares"
	"project/e-commerce/utils/helper"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type productHandler struct {
	srv product.UsecaseInterface
}

func New(e *echo.Echo, srv product.UsecaseInterface) {
	handler := productHandler{srv: srv}
	e.POST("/products", handler.Create(), middleware.JWT([]byte(cf.SECRET_JWT)))
	e.PUT("/products/:id", handler.Update(), middleware.JWT([]byte(cf.SECRET_JWT)))
	e.DELETE("/products/:id", handler.Delete(), middleware.JWT([]byte(cf.SECRET_JWT)))
	e.GET("/products", handler.ShowAll())
	e.GET("/products/:id", handler.ShowByID())
	e.GET("/myproducts", handler.ShowMy(), middleware.JWT([]byte(cf.SECRET_JWT)))
}

func (ph *productHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input ProductFormat
		userId := middlewares.ExtractToken(c)
		input.UserID = uint(userId)

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, failResponse("cannot bind input"))
		}

		file, _ := c.FormFile("images")
		if file != nil {
			res, err := helper.UploadProfile(c)
			if err != nil {
				return err
			}
			input.Images = res
		}
		// if errFile != nil {
		// 	return errFile
		// }

		cnv := ToCore(input)
		res, err := ph.srv.Create(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, failResponse(err.Error()))
		}
		return c.JSON(http.StatusCreated, successResponse("success create new product", toResponse(res, "product")))
	}
}

func (ph *productHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		var input ProductFormat
		userId := middlewares.ExtractToken(c)
		input.UserID = uint(userId)

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, failResponse("cannot bind input"))
		}

		file, _ := c.FormFile("images")
		if file != nil {
			res, err := helper.UploadProfile(c)
			if err != nil {
				return err
			}
			input.Images = res
		}
		// if errFile != nil {
		// 	return errFile
		// }

		cnv := ToCore(input)
		res, err := ph.srv.Update(cnv, uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, failResponse(err.Error()))
		}
		return c.JSON(http.StatusCreated, successResponse("success update product", toResponse(res, "product")))
	}
}

func (ph *productHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		res, err := ph.srv.Delete(uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, failResponse(err.Error()))
		}
		return c.JSON(http.StatusAccepted, successResponse("success delete product", toResponse(res, "product")))
	}
}

func (ph *productHandler) ShowAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ph.srv.ShowAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, failResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, successResponse("success get all product", toResponse(res, "all")))
	}
}

func (ph *productHandler) ShowByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		res, err := ph.srv.ShowByID(uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, failResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, successResponse("success get product by id", toResponse(res, "product")))
	}
}

func (ph *productHandler) ShowMy() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := middlewares.ExtractToken(c)
		res, err := ph.srv.ShowMy(uint(userId))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, failResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, successResponse("success get my product", toResponse(res, "all")))
	}
}
