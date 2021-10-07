package healthcheck

import (
	"context"
	"crypto/tls"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	myResty "gitlab.com/husol/hus-echo/client/resty"
)

// ErrCheckHealth .
var ErrCheckHealth = errors.New("fail to check health")

type pgRepository struct {
	getClient   func(ctx context.Context) *gorm.DB
	restyClient *resty.Client
}

func NewPGRepository(getClient func(ctx context.Context) *gorm.DB) Repository {
	return &pgRepository{
		getClient:   getClient,
		restyClient: myResty.GetClient(),
	}
}

func (r *pgRepository) CallHealthCheck(ctx context.Context, url string) error {
	// Test resty and parse data with url = https://api.publicapis.org/entries
	// var entryResp model.EntryResponse
	//nolint:gosec
	resp, err := r.restyClient.
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).R().
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
		}).
		// SetResult(&entryResp).
		Get(url)

	if err != nil {
		return errors.Wrap(err, ErrCheckHealth.Error())
	}

	if resp.IsError() {
		return ErrCheckHealth
	}

	return nil
}
