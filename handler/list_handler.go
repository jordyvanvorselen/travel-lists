package handler

import (
	"fmt"
	"net/http"

	guuid "github.com/google/uuid"
	"github.com/jordyvanvorselen/travel-lists/internal/domain"
	"github.com/jordyvanvorselen/travel-lists/internal/service"
	template "github.com/jordyvanvorselen/travel-lists/web/templates/list"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ListHandler struct {
	Db *gorm.DB
}

func (h ListHandler) New(c echo.Context) error {
	return render(c, template.New())
}

func (h ListHandler) Create(c echo.Context) error {
	var list domain.List
	var err error

	list.UUID = guuid.New()

	if err = c.Bind(&list); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	if list, err = service.CreateList(h.Db, list); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
	}

	c.Response().Header().Set("HX-Redirect", fmt.Sprintf("/lists/%s", list.UUID.String()))

	return c.NoContent(http.StatusOK)
}

func (h ListHandler) Show(c echo.Context) error {
	var list *domain.List
	var listItems []domain.ListItem
	var err error

	uuid := c.Param("uuid")

	if list, err = service.GetListByUUID(h.Db, uuid); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
	}

	if listItems, err = service.GetListItemsByListId(h.Db, list.ID); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
	}

	return render(c, template.Show(list, listItems))
}
