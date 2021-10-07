package healthcheck

import (
	"github.com/phamtrung99/gowebbasic/repository"
	"github.com/phamtrung99/gowebbasic/repository/healthcheck"
)

type Usecase struct {
	HealthCheckRepo healthcheck.Repository
}

// New .
func New(repo *repository.Repository) IUsecase {
	return &Usecase{
		HealthCheckRepo: repo.HealthCheck,
	}
}
