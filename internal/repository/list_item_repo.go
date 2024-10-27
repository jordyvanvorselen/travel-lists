package repository

import (
	"context"

	"github.com/jordyvanvorselen/travel-lists/internal/domain"
	"github.com/jordyvanvorselen/travel-lists/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func SaveListItem(ctx context.Context, listItem domain.ListItem) (domain.ListItem, error) {
	newListItem := models.ListItem{
		ID:       listItem.Id,
		ListID:   listItem.ListId,
		Location: listItem.Location,
	}

	if err := newListItem.InsertG(ctx, boil.Infer()); err != nil {
		return domain.ListItem{}, err
	}

	return listItem, nil
}

func GetListItemsByListId(ctx context.Context, id int) ([]domain.ListItem, error) {
	var listItems []*models.ListItem
	var err error

	query := models.ListItemWhere.ListID.EQ(id)

	if listItems, err = models.ListItems(query).All(ctx, boil.GetContextDB()); err != nil {
		return []domain.ListItem{}, err
	}

	var domainListItems []domain.ListItem
	for _, listItem := range listItems {
		domainListItems = append(domainListItems, domain.ListItem{
			Id:       listItem.ID,
			ListId:   listItem.ListID,
			Location: listItem.Location,
		})
	}

	return domainListItems, nil
}
