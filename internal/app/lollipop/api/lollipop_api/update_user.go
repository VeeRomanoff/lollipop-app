package lollipop_api

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	desc "github.com/VeeRomanoff/Lollipop/internal/pb/lollipop/api"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) UpdateUser(ctx context.Context, req *desc.UpdateUserRequest) (*emptypb.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}
