package lollipop_api

import (
	"context"
	"errors"

	"github.com/VeeRomanoff/Lollipop/internal/domain"
	internal_errors "github.com/VeeRomanoff/Lollipop/internal/errors"
	desc "github.com/VeeRomanoff/Lollipop/internal/pb/lollipop/api"
)

func (i *Implementation) GetUserById(ctx context.Context, req *desc.GetUserByIDRequest) (*desc.GetUserByIDResponse, error) {
	if err := validateGetUserByIDRequest(req); err != nil {
		return nil, validationError(err)
	}

	resp, err := i.userService.GetUserById(ctx, req.GetId())
	if err != nil {
		return nil, internal_errors.HandleServiceError(err)
	}

	return &desc.GetUserByIDResponse{
		User: mapUserToProto(resp),
	}, nil
}

func mapUserToProto(user *domain.User) *desc.User {
	return &desc.User{
		Id:          user.ID,
		Name:        user.Name,
		Age:         user.Age,
		Height:      user.Height,
		Hobbies:     user.Hobbies,
		Description: user.Description,
		Email:       user.Email,
	}
}

func validateGetUserByIDRequest(req *desc.GetUserByIDRequest) error {
	if req == nil {
		return errors.New("request is nil")
	}

	if req.GetId() <= 0 {
		return errors.New("invalid id")
	}

	return nil
}
