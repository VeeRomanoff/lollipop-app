package lollipop_api

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// validationError для ошибок связанных с валидацией
func validationError(err error) error {
	return status.Error(codes.InvalidArgument, fmt.Sprintf("validation error: %v", err))
}
