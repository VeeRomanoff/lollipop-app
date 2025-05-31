package error_wrapper

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func IsCanceled(err error) bool {
	return errors.Is(err, context.Canceled) || code(err) == codes.Canceled
}

func code(err error) codes.Code {
	return status.Code(err)
}
