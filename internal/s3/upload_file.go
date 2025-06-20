package s3

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"io"
	"log"
)

func (s *MinioStore) UploadImage(ctx context.Context, bucketName, fileName string, file io.Reader) error {
	// Проверка на существование бакета
	ok, err := s.client.BucketExists(ctx, bucketName)
	if err != nil {
		log.Fatalf("failed to check if bucket exists: %v", err)
	}
	if !ok {
		return fmt.Errorf("bucket does not exist")
	}
	log.Printf("uploading file to bucket %s\n", bucketName)
	_, err = s.client.PutObject(ctx, bucketName, fileName, file, -1, minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})
	if err != nil {
		return fmt.Errorf("failed to upload object: %v", err)
	}
	log.Printf("uploaded file to bucket %s\n", bucketName)

	return nil
}
