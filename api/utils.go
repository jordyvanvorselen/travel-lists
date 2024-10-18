package api

import (
	"github.com/a-h/templ"
	"github.com/jordyvanvorselen/travel-lists/web/templates/layout"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component) error {
	return layout.Layout(component).Render(c.Request().Context(), c.Response())
}
