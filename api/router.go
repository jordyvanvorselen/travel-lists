package api

import (
	"net/http"

	"github.com/jordyvanvorselen/travel-lists/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/assets", "assets")

	e.GET("/", controller.HomeHandler{}.Index)
	e.GET("/create-list", controller.ListHandler{}.New)
	e.GET("/lists", controller.ListHandler{}.Index)
	e.POST("/lists", controller.ListHandler{}.Create)

	e.ServeHTTP(w, r)
}
