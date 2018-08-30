package config

import (
	"flag"
	"errors"
	"os"
	"fmt"
)

var (
	NoUploadKey = errors.New("need to specify upload key; see -h")
)

type Config interface {
	Verbose() bool
	Tracing() bool

	UploadKey() string
	Region() string
	Bucket() string
	AwsAccessKeyId() string
	AwsSecretAccessKey() string
}

type config struct {
	verbose bool
	tracing bool

	uploadKey          string
	region             string
	bucket             string
	awsAccessKeyId     string
	awsSecretAccessKey string
}

func Init() (Config, error) {
	uploadKey := flag.String("upload_key", "", "Key for upload")
	verbose := flag.Bool("verbose", false, "Verbose mode")
	tracing := flag.Bool("trace", false, "Collect trace information")
	flag.Parse()
	if *uploadKey == "" {
		return nil, NoUploadKey
	}
	awsAccessKeyId, err := getEnv("AWS_ACCESS_KEY_ID")
	if err != nil {
		return nil, err
	}
	awsSecretAccessKey, err := getEnv("AWS_SECRET_ACCESS_KEY")
	if err != nil {
		return nil, err
	}
	region, err := getEnv("REGION")
	if err != nil {
		return nil, err
	}
	bucket, err := getEnv("BUCKET")
	if err != nil {
		return nil, err
	}
	return &config{
		verbose:            *verbose,
		uploadKey:          *uploadKey,
		tracing:            *tracing,
		awsAccessKeyId:     awsAccessKeyId,
		awsSecretAccessKey: awsSecretAccessKey,
		region:             region,
		bucket:             bucket,
	}, nil
}

func getEnv(env string) (string, error) {
	if v := os.Getenv(env); v != "" {
		return v, nil
	}
	return "", errors.New(fmt.Sprintf("Missing env %s", env))
}

func (c *config) Verbose() bool {
	return c.verbose
}

func (c *config) Tracing() bool {
	return c.tracing
}

func (c *config) UploadKey() string {
	return c.uploadKey
}

func (c *config) Region() string {
	return c.region
}

func (c *config) Bucket() string {
	return c.bucket
}

func (c *config) AwsAccessKeyId() string {
	return c.awsAccessKeyId
}

func (c *config) AwsSecretAccessKey() string {
	return c.awsSecretAccessKey
}
