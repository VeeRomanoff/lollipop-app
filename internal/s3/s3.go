package s3

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

// TODO export sensitive variables to environment config
const (
	Endpoint        = "localhost:9000"
	AccessKeyID     = "minioadmin"
	AccessKeySecret = "minioadmin"
)

type Config struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	UseSSL    bool
}

type MinioStore struct {
	Client *minio.Client
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
		Client: minioClient,
	}

	return minioStore
}
