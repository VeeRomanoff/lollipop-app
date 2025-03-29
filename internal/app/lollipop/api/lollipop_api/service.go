package lollipop_api

import (
	"context"
	lollipop "github.com/VeeRomanoff/Lollipop/internal/pb/lollipop/api"

	"github.com/VeeRomanoff/Lollipop/internal/domain"
)

// userService интерфейс для сервиса пользователей
type userService interface {
	GetUser(ctx context.Context, userID int64) (*domain.User, error)
	RegisterUser(ctx context.Context, user *domain.User) (int64, error)
	UpdateUser(ctx context.Context, user *domain.User) error
	DeleteUser(ctx context.Context, userID int64) error
}

type Implementation struct {
	lollipop.UnimplementedLollipopServer

	userService userService
}

func (i *Implementation) MustEmbedUnimplementedLollipopServer() {}

func NewLollipop(
	userService userService,
) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
