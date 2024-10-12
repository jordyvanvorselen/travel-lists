package api

import (
	"github.com/jordyvanvorselen/go-templ-htmx-vercel-template/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/lists", controller.ListHandler{}.Index)

	e.ServeHTTP(w, r)
}
