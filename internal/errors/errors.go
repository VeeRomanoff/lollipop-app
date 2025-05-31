package errors

import (
	"errors"
	"github.com/VeeRomanoff/Lollipop/internal/errors/error_wrapper"
	"google.golang.org/grpc/codes"
)

var (
	// ErrNotFound ...
	ErrNotFound = errors.New("not found")
	// ErrAlreadyExists ...
	ErrAlreadyExists = errors.New("already exists")
	// ErrInvalidArgument ...
	ErrInvalidArgument = errors.New("invalid argument")
	// ErrFailedPrecondition ...
	ErrFailedPrecondition = errors.New("failed precondition")
)

func HandleServiceError(err error) error {
	if errors.Is(err, ErrNotFound) {
		return error_wrapper.WithCode(codes.NotFound, err)
	}

	if errors.Is(err, ErrAlreadyExists) {
		return error_wrapper.WithCode(codes.AlreadyExists, err)
	}

	if errors.Is(err, ErrInvalidArgument) {
		return error_wrapper.WithCode(codes.InvalidArgument, err)
	}

	if errors.Is(err, ErrFailedPrecondition) {
		return error_wrapper.WithCode(codes.FailedPrecondition, err)
	}

	return error_wrapper.WithCode(codes.Internal, err)
}
