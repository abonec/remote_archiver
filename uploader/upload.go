package uploader

import (
	"fmt"
	"github.com/abonec/file_downloader/config"
	"github.com/abonec/file_downloader/log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"os"
)

func Upload(body io.Reader, cfg config.Config) {
	err := upload(
		body, cfg.UploadKey(),
		cfg.AwsAccessKeyId(), cfg.AwsSecretAccessKey(),
		cfg.Region(), cfg.Bucket(),
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
	uploader.PartSize = 20 * 1024 * 1024
	resp, err := uploader.Upload(upParams)

	if err != nil {
		return err
	}
	fmt.Printf("\nresponse %s\n", awsutil.StringValue(resp))
	return nil
}
