package lollipop_api

import (
	"context"
	"errors"
	"log"

	"github.com/VeeRomanoff/Lollipop/internal/domain"
	internal_errors "github.com/VeeRomanoff/Lollipop/internal/errors"
	desc "github.com/VeeRomanoff/Lollipop/internal/pb/lollipop/api"
)

func (i *Implementation) UpdateUser(ctx context.Context, req *desc.UpdateUserRequest) (*desc.User, error) {
	if err := validateUpdateUserRequest(req); err != nil {
		return nil, validationError(err)
	}

	user, err := i.userService.UpdateUser(ctx, &domain.User{
		ID:          req.GetId(),
		Name:        req.GetName(),
		Age:         req.GetAge(),
		Height:      req.GetHeight(),
		Hobbies:     req.GetHobbies(),
		Email:       req.GetEmail(),
		Description: req.GetDescription(),
	})
	log.Printf("email: %v", req.GetEmail())
	if err != nil {
		return nil, internal_errors.HandleServiceError(err)
	}

	return &desc.User{
		Id:          user.ID,
		Name:        user.Name,
		Age:         user.Age,
		Hobbies:     user.Hobbies,
		Height:      user.Height,
		Description: user.Description,
		Email:       user.Email,
	}, nil
}

func validateUpdateUserRequest(req *desc.UpdateUserRequest) error {
	if req == nil {
		return errors.New("empty request")
	}

	if req.GetId() <= 0 {
		return errors.New("invalid id")
	}

	return nil
}
