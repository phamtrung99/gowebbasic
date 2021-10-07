package userfavorite

import (
	"context"

	"github.com/phamtrung99/gowebbasic/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type pgRepository struct {
	getClient func(ctx context.Context) *gorm.DB
}

func NewPGRepository(getClient func(ctx context.Context) *gorm.DB) Repository {
	return &pgRepository{getClient}
}

func (r *pgRepository) Create(ctx context.Context, userfavor *model.UserFavorite) (*model.UserFavorite, error) {
	db := r.getClient(ctx)
	err := db.Create(userfavor).Error

	return userfavor, errors.Wrap(err, "create user favorite")
}

func (r *pgRepository) GetListMovieIDByUserID(ctx context.Context, userID int64) ([]int64, error) {
	db := r.getClient(ctx)
	var userFavor []model.UserFavorite
	var result = []int64{}

	err := db.Where("user_id = ? ", userID).
		Find(&userFavor).Error

	if err != nil {
		return result, errors.Wrap(err, "find user favorite movie")
	}

	for i := 0; i < len(userFavor); i++ {
		result = append(result, userFavor[i].MovieID)
	}

	return result, nil
}

func (r *pgRepository) GetByIDMovieAndIDUser(ctx context.Context, userID int64, movieID int64) (*model.UserFavorite, error) {
	db := r.getClient(ctx)
	userFavor := &model.UserFavorite{}

	err := db.Where("user_id = ? AND movie_id = ?", userID, movieID).First(userFavor).Error

	if err != nil {
		return &model.UserFavorite{}, errors.Wrap(err, "get user favorite by id movie and id user")
	}

	return userFavor, nil
}
