package controller

import (
	"github.com/jordyvanvorselen/go-templ-htmx-vercel-template/view/list"
	"github.com/labstack/echo/v4"
)

type ListHandler struct{}

func (h ListHandler) Index(c echo.Context) error {
	return render(c, list.Index())
}
