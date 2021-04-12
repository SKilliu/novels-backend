package s3

import (
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/sirupsen/logrus"
)

var s3Client *Client

type Client struct {
	minio  *minio.Client
	bucket string
}

func Init(logger *logrus.Entry) {
	var err error

	setLogger(logger)

	loadConfigFromEnvs()

	s3Client, err = New(configuration)
	if err != nil {
		panic(err)
	}
}

func New(config Configuration) (*Client, error) {
	client, err := minio.New(config.Url, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKey, config.SecretKey, ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}

	return &Client{
		minio:  client,
		bucket: config.Bucket,
	}, nil
}

func (c Client) URL(fileName string) string {
	return fmt.Sprintf("http://localhost:9000/%s/%s", c.bucket, fileName)
}

func S3Client() *Client {
	return s3Client
}
