package api

import (
	"net/http"

	"github.com/jordyvanvorselen/travel-lists/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/assets", "web/assets")

	e.GET("/", handlers.HomeHandler{}.Index)
	e.GET("/create-list", handlers.ListHandler{}.New)
	e.GET("/lists/:id", handlers.ListHandler{}.Show)
	e.POST("/lists", handlers.ListHandler{}.Create)

	e.ServeHTTP(w, r)
}
