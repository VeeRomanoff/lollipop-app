package users_service

import (
	"context"
	"fmt"
	"log"

	"github.com/VeeRomanoff/Lollipop/internal/database"
	"github.com/VeeRomanoff/Lollipop/internal/domain"
	internal_errors "github.com/VeeRomanoff/Lollipop/internal/errors"
)

var id int64

type Service struct {
	DB *database.Database
}

// RegisterUser регистрация юзера
func (s *Service) RegisterUser(ctx context.Context, user *domain.User) (int64, error) {
	user.ID = id + 1

	// TODO check user's existance
	// TODO worrk with errors

	id, err := s.DB.RegisterUser(ctx, user)
	if err != nil {
		return 0, fmt.Errorf("%w", err)
	}
	return id, nil
}

// GetUserById получение юзера по ID
func (s *Service) GetUserById(ctx context.Context, userID int64) (*domain.User, error) {
	user, err := s.DB.GetUserByID(ctx, userID)
	if user == nil {
		return nil, internal_errors.ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("getting user by id: %w", err)
	}
	log.Printf("user: %v", user)

	return user, nil
}

// UpdateUser обновление юзера
func (s *Service) UpdateUser(ctx context.Context, userReq *domain.User) (*domain.User, error) {
	userToUpdate, err := s.validateUpdateUser(ctx, userReq)
	if err != nil {
		return nil, err
	}

	updatedUser, err := s.DB.UpdateUser(ctx, userToUpdate)
	if err != nil {
		return nil, fmt.Errorf("updating user: %w", err)
	}

	return updatedUser, nil
}

func (s *Service) validateUpdateUser(ctx context.Context, userReq *domain.User) (*domain.User, error) {
	if userReq == nil {
		return nil, internal_errors.ErrInvalidArgument
	}

	// Получаем пользователя по id
	userExists, err := s.GetUserById(ctx, userReq.ID)
	if userExists == nil {
		return nil, internal_errors.ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("getting user by id: %w", err)
	}

	// TODO сначала проверить с новой структурой вс полями от существующегшо юзера
	// todo потом проверить сам существующий юзер и на нем обновлять поля

	// Обновляемый пользователь
	var userToUpdate = &domain.User{
		ID:          userExists.ID,
		Name:        userExists.Name,
		Age:         userExists.Age,
		Height:      userExists.Height,
		Description: userExists.Description,
		Hobbies:     userExists.Hobbies,
		Email:       userExists.Email,
	}

	hasUpdates := false

	if userReq.Name != "" {
		userToUpdate.Name = userReq.Name
		hasUpdates = true
	}

	if userReq.Age != 0 {
		if userReq.Age <= 18 {
			log.Println("userReq.Age is less than 18")
			return nil, internal_errors.ErrInvalidArgument
		}
		userToUpdate.Age = userReq.Age
		hasUpdates = true
	}

	if userReq.Description != "" {
		userToUpdate.Description = userReq.Description
		hasUpdates = true
	}

	if userReq.Email != "" {
		log.Println("in email...")
		userToUpdate.Email = userReq.Email
		hasUpdates = true
	}

	if userReq.Height != 0 {
		userToUpdate.Height = userReq.Height
		hasUpdates = true
	}

	if userReq.Hobbies != nil {
		userToUpdate.Hobbies = userReq.Hobbies
		hasUpdates = true
	}

	if !hasUpdates {
		log.Println("has updates has triggered")
		return nil, internal_errors.ErrInvalidArgument
	}
	log.Println("all passed")
	return userToUpdate, nil
}

// DeleteUser удаление юзера
func (s *Service) DeleteUser(ctx context.Context, userID int64) error {
	if err := s.DB.DeleteUser(ctx, userID); err != nil {
		return err
	}

	return nil
}
