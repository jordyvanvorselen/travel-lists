package handlers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/jordyvanvorselen/travel-lists/internal/domain"
	"github.com/jordyvanvorselen/travel-lists/internal/service"
	template "github.com/jordyvanvorselen/travel-lists/web/templates/list"
	"github.com/labstack/echo/v4"
)

type ListHandler struct{}

func (h ListHandler) New(c echo.Context) error {
	return render(c, template.New())
}

func (h ListHandler) Create(c echo.Context) error {
	var list domain.List
	var err error

	list.UUID = uuid.New()

	if err = c.Bind(&list); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	if list, err = service.CreateList(c.Request().Context(), list); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
	}

	c.Response().Header().Set("HX-Redirect", fmt.Sprintf("/lists/%s", list.UUID.String()))

	return c.NoContent(http.StatusOK)
}

func (h ListHandler) Show(c echo.Context) error {
	var list domain.List
	var err error

	id := c.Param("id")

	if list, err = service.GetListByUUID(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
	}

	return render(c, template.Show(list))
}
