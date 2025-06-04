package httphandlers

import (
	"context"
	"io"
)

type minioStore interface {
	UploadImage(ctx context.Context, bucketName, filename string, file io.Reader) error
}

type HTTPHandler struct {
	minioStore minioStore
}

func NewHTTPHandler(minioStore minioStore) *HTTPHandler {
	return &HTTPHandler{
		minioStore: minioStore,
	}
}
