package user

import (
	"context"
	"strconv"

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

func (r *pgRepository) Delete(ctx context.Context, id int64) error {
	db := r.getClient(ctx)
	err := db.Where("ID = ?", id).
		Delete(&model.User{}).Error
	return errors.Wrap(err, "delete fail")
}

func (r *pgRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	db := r.getClient(ctx)
	err := db.Create(user).Error

	return user, errors.Wrap(err, "create user")
}

func (r *pgRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	db := r.getClient(ctx)
	user := &model.User{}

	err := db.Where("email = ?", email).
		First(user).Error

	if err != nil {
		return nil, errors.Wrap(err, "get user by email")
	}

	return user, nil
}

func (r *pgRepository) GetByID(ctx context.Context, id int64) (*model.User, error) {
	db := r.getClient(ctx)
	user := &model.User{}

	err := db.Where("id = ?", id).
		First(user).Error

	if err != nil {
		return nil, errors.Wrap(err, "get user by id")
	}

	return user, nil
}

func (r *pgRepository) Update(ctx context.Context, user *model.User) (*model.User, error) {
	db := r.getClient(ctx)

	// user, isExisted := r.GetUserByID(id)

	// if !isExisted {
	// 	return &model.User{}, false
	// }

	err := db.Model(user).
		Omit("role", "id", "deleted_at", "password", "ref_token").
		Updates(user).Error

	return user, errors.Wrap(err, "update user")
}

func (r *pgRepository) GetNextIDIncrement(ctx context.Context) string {
	db := r.getClient(ctx)
	var nextID int
	db.Raw(`
	SELECT AUTO_INCREMENT
	FROM information_schema.TABLES
	WHERE TABLE_SCHEMA = "the_movie_db"
	AND TABLE_NAME = "users"`).Scan(&nextID)

	nextID += 1

	return strconv.Itoa(nextID)
}

func (r *pgRepository) UpdateUserPassword(ctx context.Context, passwHash string, id int64) error {
	db := r.getClient(ctx)

	err := db.Where("id = ?", id).
		Updates(&model.User{Password: passwHash}).Error

	if err != nil {
		return errors.Wrap(err, "update password")
	}

	return nil
}
