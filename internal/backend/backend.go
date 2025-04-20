package backend

import (
	"context"
	"errors"

	"github.com/hashicorp/vault/sdk/logical"
)

var errUnimplemented = errors.New("unimplemented")

func Factory(ctx context.Context, conf *logical.BackendConfig) (logical.Backend, error) {
	return nil, errUnimplemented
}
