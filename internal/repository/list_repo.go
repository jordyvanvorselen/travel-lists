package repository

import (
	"context"

	"github.com/jordyvanvorselen/travel-lists/internal/domain"
	"github.com/jordyvanvorselen/travel-lists/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func Save(ctx context.Context, list domain.List) (domain.List, error) {
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
