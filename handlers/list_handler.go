package handlers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/jordyvanvorselen/travel-lists/internal/domain"
	"github.com/jordyvanvorselen/travel-lists/internal/service"
	"github.com/jordyvanvorselen/travel-lists/web/templates/list"
	"github.com/labstack/echo/v4"
)

type ListHandler struct{}

func (h ListHandler) New(c echo.Context) error {
	return render(c, list.New())
}

func (h ListHandler) Create(c echo.Context) error {
	newList := domain.List{
		UUID: uuid.New(),
	}

	err := c.Bind(&newList)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	newList, err = service.CreateList(c.Request().Context(), newList)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
	}

	c.Response().Header().Set("HX-Redirect", fmt.Sprintf("/lists/%s", newList.UUID.String()))

	return c.NoContent(http.StatusOK)
}

func (h ListHandler) Show(c echo.Context) error {
	id := c.Param("id")

	return render(c, list.Show(id))
}
