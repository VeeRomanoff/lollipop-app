package error_wrapper

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Error struct {
	code codes.Code
	err  error
}

func (e Error) Error() string {
	return e.err.Error()
}

func New(code codes.Code, msg string) error {
	return Error{
		code: code,
		err:  errors.New(msg),
	}
}

func WithCode(code codes.Code, err error) error {
	if err == nil {
		return nil
	}
	if IsCanceled(err) {
		code = codes.Canceled
	} else if s, ok := status.FromError(err); ok {
		code = s.Code()
	}

	return Error{
		code: code,
		err:  err,
	}
}

func (e Error) GRPCStatus() *status.Status {
	return status.New(e.code, e.Error())
}
