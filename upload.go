package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
	"github.com/dustin/go-humanize"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func runUpload(file, key string) {
	Trace.Println("Uploading started")
	err := upload(
		file, key,
		os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"),
		os.Getenv("REGION"), os.Getenv("BUCKET"),
	)
	if logError(err) {
		os.Exit(1)
	}
	Trace.Println("Uploading finished")
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
	if logError(err) {
		os.Exit(1)
	}
	stats, err := file.Stat()
	if logError(err) {
		os.Exit(1)
	}
	Trace.Printf("Upload file size is %s", humanize.Bytes(uint64(stats.Size())))
	if err != nil {
		return err
	}
	defer file.Close()
	//params := &s3.PutObjectInput{
	//	Bucket: aws.String(bucket),
	//	Key:    aws.String(key),
	//	Body:   file,
	//}
	//resp, err := svc.PutObject(params)
	upParams := &s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   file,
	}
	uploader := s3manager.NewUploaderWithClient(svc)
	resp, err := uploader.Upload(upParams)

	if err != nil {
		return err
	}
	fmt.Printf("response %s\n", awsutil.StringValue(resp))
	return nil
}
