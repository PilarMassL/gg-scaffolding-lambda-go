package adapter

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

const (
	AWS_S3_REGION = "us-west-2" // Region
	AWS_S3_BUCKET = "test"      // Bucket
)

type MyStorageUsingS3 struct {
	awsS3Client *s3.Client
}

func (s *MyStorageUsingS3) NewMyStorageUsingS3() *MyStorageUsingS3 {
	return &MyStorageUsingS3{
		awsS3Client: configS3(),
	}
}

func (s *MyStorageUsingS3) PutFile(key string, file *os.File, contentType, perm string) error {
	uploader := manager.NewUploader(s.awsS3Client)
	_, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(AWS_S3_BUCKET),
		Key:    aws.String(key),
		Body:   file,
	})
	return err
}

func (s *MyStorageUsingS3) GetFile(key, version string) (*os.File, error) {
	downloader := manager.NewDownloader(s.awsS3Client)
	// Create a file to write the S3 Object contents to.
	f, err := os.Create(key)
	if err != nil {
		return nil, err
	}

	// :bulb: si requiere el número de bytes descargados agregue la variable numBytes
	numBytes, err := downloader.Download(context.TODO(), f, &s3.GetObjectInput{
		Bucket: aws.String(AWS_S3_BUCKET),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	log.Printf("Se recibieron %d bytes", numBytes)
	return f, nil
}

func configS3() *s3.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(AWS_S3_REGION))
	if err != nil {
		log.Fatal(err)
	}

	return s3.NewFromConfig(cfg)
}
