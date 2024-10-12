package api

import (
	"github.com/jordyvanvorselen/travel-lists/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", controller.ListHandler{}.Index)
	e.GET("/lists", controller.ListHandler{}.Index)

	e.ServeHTTP(w, r)
}
