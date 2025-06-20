package lollipop_api

import (
	"context"
	"errors"

	internal_errors "github.com/VeeRomanoff/Lollipop/internal/errors"
	desc "github.com/VeeRomanoff/Lollipop/internal/pb/lollipop/api"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) DeleteUser(ctx context.Context, req *desc.DeleteUserRequest) (*emptypb.Empty, error) {
	if err := validateDeleteUser(req); err != nil {
		return &emptypb.Empty{}, validationError(err)
	}

	if err := i.userService.DeleteUser(ctx, req.GetUserId()); err != nil {
		return &emptypb.Empty{}, internal_errors.HandleServiceError(err)
	}

	return &emptypb.Empty{}, nil
}

func validateDeleteUser(req *desc.DeleteUserRequest) error {
	if req.GetUserId() == 0 {
		return errors.New("invalid id")
	}

	return nil
}
