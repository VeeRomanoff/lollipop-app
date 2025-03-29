package lollipop_api

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/VeeRomanoff/Lollipop/internal/pb/lollipop/api"
)

func (i *Implementation) DeleteUser(ctx context.Context, req *desc.DeleteUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented")
}
