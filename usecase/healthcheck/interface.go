package healthcheck

import (
	"context"
)

// IUsecase .
type IUsecase interface {
	CheckHealth(ctx context.Context, url string) error
}
