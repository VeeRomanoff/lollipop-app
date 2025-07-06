package users_service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/VeeRomanoff/Lollipop/internal/domain"
)

func (s *Service) GetUserByIdFromCache(ctx context.Context, userID int64) (*domain.User, error) {
	bytes, err := s.redis.Get(ctx, fmt.Sprintf("user:%d", userID)).Bytes()
	if err != nil {
		return nil, fmt.Errorf("error redis get %d, %w", userID, err)
	}

	var user *domain.User
	if err = json.Unmarshal(bytes, &user); err != nil {
		return nil, fmt.Errorf("error unmarshaling user: %w", err)
	}

	return user, nil
}
