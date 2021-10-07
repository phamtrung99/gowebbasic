package healthcheck

import (
	"context"

	"github.com/pkg/errors"
)

func (u Usecase) CheckHealth(ctx context.Context, url string) error {
	if url == "" {
		return nil
	}

	if err := u.HealthCheckRepo.CallHealthCheck(ctx, url); err != nil {
		return errors.Wrap(err, "failed to call request to health check")
	}

	return nil
}
