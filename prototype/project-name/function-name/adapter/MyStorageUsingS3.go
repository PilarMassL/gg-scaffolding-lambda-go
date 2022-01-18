package adapter

import (
	"bytes"
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

const (
	AWS_S3_REGION = "us-east-2"   // Region
	AWS_S3_BUCKET = "test-web-s3" // Bucket
)

type MyStorageUsingS3 struct {
	awsS3Client *s3.Client
}

func NewMyStorageUsingS3() *MyStorageUsingS3 {
	return &MyStorageUsingS3{
		awsS3Client: configS3(),
	}
}

func (s *MyStorageUsingS3) PutFile(key string, content []byte) error {
	r := bytes.NewReader(content)
	uploader := manager.NewUploader(s.awsS3Client)
	_, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(AWS_S3_BUCKET),
		Key:    aws.String(key),
		Body:   r,
	})
	return err
}

func (s *MyStorageUsingS3) GetFile(key, version string) ([]byte, error) {
	downloader := manager.NewDownloader(s.awsS3Client)

	awsBucket := aws.String(AWS_S3_BUCKET)
	awsKey := aws.String(key)

	// Intentamos obtener el HeadObject para saber el tamaño del archivo a cargar en memoria.
	headObject, errHeadObject := s.awsS3Client.HeadObject(context.TODO(), &s3.HeadObjectInput{
		Bucket: awsBucket,
		Key:    awsKey,
	})
	if errHeadObject != nil {
		return nil, errHeadObject
	}

	buf := make([]byte, int(headObject.ContentLength))
	// wrap with aws.WriteAtBuffer
	w := manager.NewWriteAtBuffer(buf)
	// download file into the memory

	numBytes, err := downloader.Download(context.TODO(), w, &s3.GetObjectInput{
		Bucket: awsBucket,
		Key:    awsKey,
	})
	if err != nil {
		return nil, err
	}
	log.Printf("Se recibieron %d bytes", numBytes)
	return w.Bytes(), nil
}

func configS3() *s3.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(AWS_S3_REGION))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return s3.NewFromConfig(cfg)
}
