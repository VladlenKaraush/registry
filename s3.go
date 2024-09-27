package main

import (
	"bytes"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	s3Conn *s3.Client
}

func GetS3Client(config Config) S3Client {
	const defaultRegion = "us-east-1"
	fmt.Printf("url = %s, usrname = %s, pass = %s\n", config.S3.Url, config.S3.Username, config.S3.Password)
	staticResolver := aws.EndpointResolverWithOptionsFunc(
		func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				PartitionID:       "aws",
				URL:               config.S3.Url,
				SigningRegion:     defaultRegion,
				HostnameImmutable: true,
			}, nil
		})
	cfg := aws.Config{
		Region:                      defaultRegion,
		Credentials:                 credentials.NewStaticCredentialsProvider(config.S3.Username, config.S3.Password, ""),
		EndpointResolverWithOptions: staticResolver,
	}

	s3Conn := s3.NewFromConfig(cfg)
	return S3Client{s3Conn: s3Conn}
}

func (cl *S3Client) UploadPackage(body []byte, bucket, key string) {
	cl.s3Conn.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(body),
	})
}
