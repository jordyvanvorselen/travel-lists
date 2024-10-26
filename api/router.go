package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/jordyvanvorselen/travel-lists/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	e := echo.New()

	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s port=5432 dbname=travel-lists sslmode=%s",
		os.Getenv("PSQL_HOST"), os.Getenv("PSQL_USER"), os.Getenv("PSQL_PASS"), os.Getenv("PSQL_SSLMODE"),
	)

	db, err := sql.Open("postgres", connStr)
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
