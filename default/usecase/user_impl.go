package usecase

import (
	"context"

	"github.com/abyssparanoia/catharsis-gcp/default/domain/model"
	"github.com/abyssparanoia/catharsis-gcp/default/domain/repository"
	"github.com/abyssparanoia/catharsis-gcp/internal/pkg/log"
)

type user struct {
	userRepo repository.User
}

func (s *user) Get(ctx context.Context, userID string) (*model.User, error) {
	user, err := s.userRepo.Get(ctx, userID)
	if err != nil {
		log.Errorm(ctx, "s.userRepo.Get", err)
		return nil, err
	}

	if !user.IsExist() {
		return nil, newUserNotExistError(ctx, userID)
	}

	return user, nil
}

// NewUser ... get User usecase
func NewUser(userRepo repository.User) User {
	return &user{userRepo}
}
