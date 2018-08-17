package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

func runUpload(file, key string) {
	err := upload(
		file, key,
		os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"),
		os.Getenv("REGION"), os.Getenv("BUCKET"),
	)
	if logError(err) {
		os.Exit(1)
	}
}

func upload(filePath, key, awsAccessKeyId, awsSecretAccessKey, region, bucket string) error {
	creds := credentials.NewStaticCredentials(awsAccessKeyId, awsSecretAccessKey, "")
	_, err := creds.Get()
	if err != nil {
		return err
	}
	cfg := aws.NewConfig().WithRegion(region).WithCredentials(creds)
	svc := s3.New(session.New(), cfg)

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	params := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   file,
	}
	resp, err := svc.PutObject(params)
	if err != nil {
		return err
	}
	fmt.Printf("response %s", awsutil.StringValue(resp))
	return nil
}
