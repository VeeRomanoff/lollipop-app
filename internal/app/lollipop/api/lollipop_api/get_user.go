package lollipop_api

import (
	"context"
	desc "github.com/VeeRomanoff/Lollipop/internal/pb/lollipop/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetUser(ctx context.Context, req *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented")
}
