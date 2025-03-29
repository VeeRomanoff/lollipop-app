package users_service

import (
	"context"
	"fmt"
	"github.com/VeeRomanoff/Lollipop/internal/database"
	"github.com/VeeRomanoff/Lollipop/internal/domain"
)

var id int64

type Service struct {
	DB *database.Database
}

// RegisterUser регистрация юзера
func (s *Service) RegisterUser(ctx context.Context, user *domain.User) (int64, error) {
	user.ID = id + 1
	id, err := s.DB.RegisterUser(ctx, user)
	if err != nil {
		return 0, fmt.Errorf("%w", err)
	}
	return id, nil
}

// GetUser получение юзера по ID
func (s *Service) GetUser(ctx context.Context, userID int64) (*domain.User, error) {
	return nil, nil
}

// UpdateUser обновление юзера
func (s *Service) UpdateUser(ctx context.Context, user *domain.User) error {
	return nil
}

// DeleteUser удаление юзера
func (s *Service) DeleteUser(ctx context.Context, userID int64) error {
	return nil
}
