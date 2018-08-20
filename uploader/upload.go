package uploader

import (
	"os"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"io"
	"github.com/abonec/file_downloader/log"
)

func Upload(body io.Reader, uploadKey string, verbose bool) {
	err := upload(
		body, uploadKey,
		os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"),
		os.Getenv("REGION"), os.Getenv("BUCKET"),
	)
	if log.Error(err) {
		os.Exit(1)
	}

}

func upload(body io.Reader, key, awsAccessKeyId, awsSecretAccessKey, region, bucket string) error {
	creds := credentials.NewStaticCredentials(awsAccessKeyId, awsSecretAccessKey, "")
	_, err := creds.Get()
	if err != nil {
		return err
	}
	cfg := aws.NewConfig().WithRegion(region).WithCredentials(creds)
	sess, err := session.NewSession()
	if err != nil {
		return err
	}
	svc := s3.New(sess, cfg)

	upParams := &s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   body,
	}
	uploader := s3manager.NewUploaderWithClient(svc)
	resp, err := uploader.Upload(upParams)

	if err != nil {
		return err
	}
	fmt.Printf("response %s\n", awsutil.StringValue(resp))
	return nil
}
