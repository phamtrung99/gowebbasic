package auth

import (
	"context"

	"github.com/phamtrung99/gowebbasic/model"
)

// IUsecase .
type IUsecase interface {
	Login(ctx context.Context, req LoginRequest) (string, error) 
	Register(ctx context.Context, req RegisterRequest) (*model.User, error)
}
