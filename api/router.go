package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/assets", "web/assets")

	e.GET("/", HomeHandler{}.Index)
	e.GET("/create-list", ListHandler{}.New)
	e.GET("/lists/:id", ListHandler{}.Show)
	e.POST("/lists", ListHandler{}.Create)

	e.ServeHTTP(w, r)
}
