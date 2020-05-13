package main

import (
	"erp-go-v2/app/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type M map[string]interface{}

func main() {
	r := echo.New()
	r.GET("/company", func(ctx echo.Context) error {
		var company = new(model.Company)
		data, err := company.Read()
		if err != nil {
			fmt.Println(err)
			return err
		}
		return ctx.JSON(http.StatusOK, data)
	})
	r.GET("/company/:id", func(ctx echo.Context) error {
		var company = new(model.Company)
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			return err
		}
		data, err := company.ReadOne(id)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return ctx.JSON(http.StatusOK, data)
	})

	r.POST("/company", func(ctx echo.Context) error {

		var company = new(model.Company)

		if err := ctx.Bind(company); err != nil {
			return err
		}
		company.Create()
		return ctx.JSON(http.StatusOK, company)
	})
	r.PUT("/company/:id", func(ctx echo.Context) error {

		var company = new(model.Company)
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			return err
		}
		if err := ctx.Bind(company); err != nil {
			return err
		}
		company.Update(id)
		return ctx.JSON(http.StatusOK, company)
	})

	r.DELETE("/company/:id", func(ctx echo.Context) error {

		var company = new(model.Company)
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			return err
		}
		if err := ctx.Bind(company); err != nil {
			return err
		}
		company.Delete(id)
		return ctx.JSON(http.StatusOK, company)
	})

	r.GET("/product", func(ctx echo.Context) error {
		var product = new(model.Product)
		data, err := product.Read()
		if err != nil {
			fmt.Println(err)
			return err
		}
		return ctx.JSON(http.StatusOK, data)
	})
	r.GET("/product/:id", func(ctx echo.Context) error {
		var product = new(model.Product)
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			return err
		}
		data, err := product.ReadOne(id)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return ctx.JSON(http.StatusOK, data)
	})

	r.POST("/product", func(ctx echo.Context) error {

		var product = new(model.Product)

		if err := ctx.Bind(product); err != nil {
			return err
		}
		product.Create()
		return ctx.JSON(http.StatusOK, product)
	})
	r.PUT("/product/:id", func(ctx echo.Context) error {

		var product = new(model.Product)
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			return err
		}
		if err := ctx.Bind(product); err != nil {
			return err
		}
		product.Update(id)
		return ctx.JSON(http.StatusOK, product)
	})

	r.DELETE("/product/:id", func(ctx echo.Context) error {

		var product = new(model.Product)
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			return err
		}
		if err := ctx.Bind(product); err != nil {
			return err
		}
		product.Delete(id)
		return ctx.JSON(http.StatusOK, product)
	})

	r.GET("/vendor", func(ctx echo.Context) error {
		var vendor = new(model.Vendor)
		data, err := vendor.Read()
		if err != nil {
			fmt.Println(err)
			return err
		}
		return ctx.JSON(http.StatusOK, data)
	})
	r.GET("/vendor/:id", func(ctx echo.Context) error {
		var vendor = new(model.Vendor)
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			return err
		}
		data, err := vendor.ReadOne(id)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return ctx.JSON(http.StatusOK, data)
	})

	r.POST("/vendor", func(ctx echo.Context) error {

		var vendor = new(model.Vendor)

		if err := ctx.Bind(vendor); err != nil {
			return err
		}
		vendor.Create()
		return ctx.JSON(http.StatusOK, vendor)
	})
	r.PUT("/vendor/:id", func(ctx echo.Context) error {

		var vendor = new(model.Vendor)
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			return err
		}
		if err := ctx.Bind(vendor); err != nil {
			return err
		}
		vendor.Update(id)
		return ctx.JSON(http.StatusOK, vendor)
	})

	r.DELETE("/vendor/:id", func(ctx echo.Context) error {

		var vendor = new(model.Vendor)
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			return err
		}
		if err := ctx.Bind(vendor); err != nil {
			return err
		}
		vendor.Delete(id)
		return ctx.JSON(http.StatusOK, vendor)
	})
	fmt.Println("Server Started at localhost :9000")
	r.Start(":9002")
}
