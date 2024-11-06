package api

import (
	"net/http"

	"github.com/jordyvanvorselen/travel-lists/database"
	"github.com/jordyvanvorselen/travel-lists/handler"
	"github.com/jordyvanvorselen/travel-lists/internal/domain"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

var db *gorm.DB

func Handler(w http.ResponseWriter, r *http.Request) {
	e := echo.New()

	if db == nil {
		db = database.Connect()
	}

	db.AutoMigrate(&domain.List{}, &domain.ListItem{})

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/web/assets", "web/assets")

	e.GET("/", handler.HomeHandler{}.Index)
	e.GET("/create-list", handler.ListHandler{Db: db}.New)
	e.GET("/lists/:uuid", handler.ListHandler{Db: db}.Show)

	e.POST("/lists", handler.ListHandler{Db: db}.Create)
	e.POST("/lists/:uuid/list-items", handler.ListItemHandler{Db: db}.Create)

	e.ServeHTTP(w, r)
}
