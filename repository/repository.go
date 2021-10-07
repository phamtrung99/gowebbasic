package repository

import (
	"context"

	"github.com/phamtrung99/gowebbasic/repository/comment"
	"github.com/phamtrung99/gowebbasic/repository/healthcheck"
	"github.com/phamtrung99/gowebbasic/repository/movie"
	"github.com/phamtrung99/gowebbasic/repository/user"
	"github.com/phamtrung99/gowebbasic/repository/userfavorite"
	"gorm.io/gorm"
)

type Repository struct {
	HealthCheck healthcheck.Repository
	User        user.Repository
	Movie       movie.Repository
	UserFavor   userfavorite.Repository
	Comment     comment.Repository
}

func New(
	getSQLClient func(ctx context.Context) *gorm.DB,
	// getRedisClient func(ctx context.Context) *redis.Client,
) *Repository {
	return &Repository{
		HealthCheck: healthcheck.NewPGRepository(getSQLClient),
		User:        user.NewPGRepository(getSQLClient),
		Movie:       movie.NewPGRepository(getSQLClient),
		UserFavor:   userfavorite.NewPGRepository(getSQLClient),
		Comment:     comment.NewPGRepository(getSQLClient),
	}
}
