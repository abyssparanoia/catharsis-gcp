package repository

import (
	"context"

	"github.com/abyssparanoia/catharsis-gcp/default/domain/model"
	"github.com/abyssparanoia/catharsis-gcp/default/domain/repository"
	"github.com/abyssparanoia/catharsis-gcp/default/infrastructure/entity"
	"github.com/abyssparanoia/catharsis-gcp/internal/pkg/log"
	"github.com/abyssparanoia/catharsis-gcp/internal/pkg/mysql"
)

type user struct {
	cli *mysql.Client
}

func (r *user) Get(ctx context.Context, userID string) (*model.User, error) {

	dsts := []*entity.User{}

	db := r.cli.GetDB(ctx).
		Where("id = ?", userID).
		Limit(1).
		Find(&dsts)

	if err := mysql.HandleErrors(db); err != nil {
		log.Errorm(ctx, "db.Find", err)
		return nil, err
	}

	if len(dsts) == 0 {
		return nil, nil
	}

	return entity.NewUsers(dsts)[0], nil
}

// NewUser ... get user repository
func NewUser(cli *mysql.Client) repository.User {
	return &user{cli}
}
