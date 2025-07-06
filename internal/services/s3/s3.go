package s3

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

type Config struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	UseSSL    bool
}

type MinioStore struct {
	client *minio.Client
}

func NewClient(config Config) *MinioStore {
	minioClient, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKey, config.SecretKey, ""),
		Secure: config.UseSSL,
	})
	if err != nil {
		log.Fatalf("Failed to create minio client: %v", err)
	}

	minioStore := &MinioStore{
		client: minioClient,
	}

	return minioStore
}
