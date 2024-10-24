package api

import (
	"database/sql"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/jordyvanvorselen/travel-lists/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	e := echo.New()

	db, err := sql.Open("postgres", "dbname=travel-lists user=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}

	boil.SetDB(db)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/web/assets", "web/assets")

	e.GET("/", handlers.HomeHandler{}.Index)
	e.GET("/create-list", handlers.ListHandler{}.New)
	e.GET("/lists/:id", handlers.ListHandler{}.Show)

	e.POST("/lists", handlers.ListHandler{}.Create)

	e.ServeHTTP(w, r)
}
