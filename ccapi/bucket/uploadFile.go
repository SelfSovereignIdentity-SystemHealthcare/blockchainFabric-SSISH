package bucket

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/minio/minio-go/v7"
)

func UploadFile(ctx context.Context, minioClient *minio.Client, filePath string) (string, error) {
	// Make a new bucket.
	bucketName := os.Getenv("MINIO_BUCKET_NAME")
	location := "us-east-1"

	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("Bucket %s already exists\n", bucketName)
		} else {
			return "", err
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	// Upload the test file
	// Change the value of filePath if the file is in another location
	splittedFilePath := strings.Split(filePath, "/")

	objectName := splittedFilePath[len(splittedFilePath)-1]
	contentType := "application/octet-stream"

	// Upload the test file with FPutObject
	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return "", err
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)

	return bucketName + "/" + objectName, nil
}
