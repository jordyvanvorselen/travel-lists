package repository

import (
	"context"

	guuid "github.com/google/uuid"
	"github.com/jordyvanvorselen/travel-lists/internal/domain"
	"github.com/jordyvanvorselen/travel-lists/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func SaveList(ctx context.Context, list domain.List) (domain.List, error) {
	newList := models.List{
		ID:   list.Id,
		Name: list.Name,
		UUID: list.UUID.String(),
	}

	if err := newList.InsertG(ctx, boil.Infer()); err != nil {
		return domain.List{}, err
	}

	return list, nil
}

func GetListByUUID(ctx context.Context, uuid string) (domain.List, error) {
	var list *models.List
	var err error
	var listUUID guuid.UUID

	query := models.ListWhere.UUID.EQ(uuid)

	if list, err = models.Lists(query).One(ctx, boil.GetContextDB()); err != nil {
		return domain.List{}, err
	}

	if listUUID, err = guuid.Parse(list.UUID); err != nil {
		return domain.List{}, err
	}

	return domain.List{Id: list.ID, Name: list.Name, UUID: listUUID}, nil
}
