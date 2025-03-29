package domain

import "errors"

var (
	// ErrorInvalidArgument ошибка INVALID_ARGUMENT
	ErrorInvalidArgument = errors.New("invalid argument")
	// ErrorNotFound ошибка NOT_FOUND
	ErrorNotFound = errors.New("not found")
	// ErrorInternal ошибка INTERNAL_SERVER_ERROR
	ErrorInternal = errors.New("internal error")
)
