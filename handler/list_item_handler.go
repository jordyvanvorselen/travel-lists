package handler

import (
	"net/http"

	"github.com/jordyvanvorselen/travel-lists/internal/domain"
	"github.com/jordyvanvorselen/travel-lists/internal/service"
	template "github.com/jordyvanvorselen/travel-lists/web/templates/list"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ListItemHandler struct {
	Db *gorm.DB
}

func (h ListItemHandler) Create(c echo.Context) error {
	var list *domain.List
	var newListItem domain.ListItem
	var listItems []domain.ListItem
	var err error

	if err = c.Bind(&newListItem); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	uuid := c.Param("uuid")
	if list, err = service.GetListByUUID(h.Db, uuid); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
	}

	newListItem.ListID = list.ID

	if newListItem, err = service.CreateListItem(h.Db, newListItem); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
	}

	if listItems, err = service.GetListItemsByListId(h.Db, list.ID); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
	}

	return render(c, template.Show(list, listItems))
}
