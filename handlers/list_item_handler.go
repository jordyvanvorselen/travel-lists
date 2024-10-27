package handlers

import (
	"net/http"

	"github.com/jordyvanvorselen/travel-lists/internal/domain"
	"github.com/jordyvanvorselen/travel-lists/internal/service"
	template "github.com/jordyvanvorselen/travel-lists/web/templates/list"
	"github.com/labstack/echo/v4"
)

type ListItemHandler struct{}

func (h ListItemHandler) Create(c echo.Context) error {
	var list domain.List
	var newListItem domain.ListItem
	var listItems []domain.ListItem
	var err error

	if err = c.Bind(&newListItem); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	uuid := c.Param("uuid")
	if list, err = service.GetListByUUID(c.Request().Context(), uuid); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
	}

	newListItem.ListId = list.Id

	if newListItem, err = service.CreateListItem(c.Request().Context(), newListItem); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
	}

	if listItems, err = service.GetListItemsByListId(c.Request().Context(), list.Id); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
	}

	return render(c, template.Show(list, listItems))
}
