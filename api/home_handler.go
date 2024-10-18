package api

import (
	"github.com/jordyvanvorselen/travel-lists/web/templates/home"
	"github.com/labstack/echo/v4"
)

type HomeHandler struct{}

func (h HomeHandler) Index(c echo.Context) error {
	return render(c, home.Home())
}
