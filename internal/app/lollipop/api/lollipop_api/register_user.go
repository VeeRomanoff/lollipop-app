package lollipop_api

import (
	"context"
	"errors"
	"github.com/VeeRomanoff/Lollipop/internal/domain"
	desc "github.com/VeeRomanoff/Lollipop/internal/pb/lollipop/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RegisterUser создание пользователя
func (i *Implementation) RegisterUser(ctx context.Context, req *desc.RegisterUserRequest) (*desc.RegisterUserResponse, error) {
	if err := validateRegisterUser(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	id, err := i.userService.RegisterUser(ctx, &domain.User{
		Name:        req.GetName(),
		Age:         req.GetAge(),
		Height:      req.GetHeight(),
		Hobbies:     req.GetHobbies(),
		Description: req.GetDescription(),
		Email:       req.GetEmail(),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &desc.RegisterUserResponse{
		UserId: id,
	}, nil
}

func validateRegisterUser(req *desc.RegisterUserRequest) error {
	if req == nil {
		return errors.New("empty request")
	}

	if req.Age < 18 {
		return errors.New("invalid age")
	}

	return nil
}
