package bucket

import (
	"context"
	"io"
	"time"

	"github.com/minio/minio-go/v7"
)

func GetObject(ctx context.Context, minioClient *minio.Client, bucketName, objectName string) ([]byte, string, error) {
	object, err := minioClient.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, "", err
	}

	objectBytes, err := io.ReadAll(object)
	if err != nil {
		return nil, "", err
	}

	expiry := time.Duration(7 * 24 * time.Hour)
	url, err := minioClient.PresignedGetObject(ctx, bucketName, objectName, expiry, nil)
	if err != nil {
		return nil, "", err
	}

	return objectBytes, url.String(), nil
}
