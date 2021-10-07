package healthcheck

import (
	"context"
)

// Repository .
type Repository interface {
	CallHealthCheck(ctx context.Context, url string) error
}
