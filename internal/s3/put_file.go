package s3

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
)

func (c Client) PutObject(in multipart.File, size int64, fNameSeed string) (string, error) {
	err := c.minio.MakeBucket(context.Background(), c.bucket, minio.MakeBucketOptions{Region: "west-europe2"})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		_, errBucketExists := c.minio.BucketExists(context.Background(), c.bucket)
		if errBucketExists != nil {
			logger.WithError(errBucketExists).Error("failed to create the bucket")
			return "", err
		}
	}

	fName := fmt.Sprintf("photo-%s.jpeg", fNameSeed)

	_, err = c.minio.PutObject(context.Background(), c.bucket, fName, in, size, minio.PutObjectOptions{ContentType: "image/jpeg"})
	if err != nil {
		logger.WithError(err).Error("failed to put file in s3")
		return "", err
	}

	return c.URL(fName), nil
}
