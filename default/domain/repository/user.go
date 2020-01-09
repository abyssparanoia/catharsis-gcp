package repository

import (
	"context"

	"github.com/abyssparanoia/catharsis-gcp/default/domain/model"
)

// User ... user interface
type User interface {
	Get(ctx context.Context, userID string) (*model.User, error)
}
