package clients

import (
	"context"
	"errors"
	"github.com/NuttayotSukkum/batch_consumer/internals/pkg/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	logger "github.com/labstack/gommon/log"
	"io"
	"os"
	"path/filepath"
)

type S3Client struct {
	svc         *s3.Client
	bucketNames string
}

func NewS3Client(ctx context.Context, region, bucketName, accessKey, secretKey string) (*S3Client, error) {
	creds := AwsClient(ctx, accessKey, secretKey)
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region), config.WithCredentialsProvider(creds))
	if err != nil {
		logger.Errorf("ctx : %v Error AWS config: %s", ctx, err)
		return nil, err
	}
	svc := s3.NewFromConfig(cfg)
	return &S3Client{
		svc:         svc,
		bucketNames: bucketName,
	}, nil
}

func (client *S3Client) DownloadFile(ctx context.Context, dir string) error {
	result, err := client.svc.ListObjects(ctx, &s3.ListObjectsInput{
		Bucket: aws.String(client.bucketNames),
	})
	if err != nil {
		logger.Errorf("ctx: %s Error listing files from S3: %s", ctx, err)
		return err
	}
	if len(result.Contents) == 0 {
		logger.Warnf("ctx: %s No files found in S3 bucket", ctx)
		return err
	}
	logger.Infof("ctx: %s Total files in S3 bucket: %d", ctx, len(result.Contents))

	totalFileDownload := 0
	for _, item := range result.Contents {
		if utils.IsFileInRange(*item.Key) {
			// Download the file
			if err := client.downloadFileFromS3(ctx, dir, &item); err != nil {
				logger.Errorf("ctx: %s Error downloading file %s: %s", ctx, *item.Key, err)
				return err
			}
			totalFileDownload++
		}
	}
	if totalFileDownload == 0 {
		logger.Errorf("ctx: %s No files matched the criteria in S3 bucket:%s", ctx, client.bucketNames)
		return errors.New("No files matched the criteria in S3 bucket")
	}
	return nil
}

func (client *S3Client) downloadFileFromS3(ctx context.Context, dir string, item *types.Object) error {
	obj, err := client.svc.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(client.bucketNames),
		Key:    aws.String(*item.Key),
	})
	if err != nil {
		logger.Errorf("dto getting file from S3: %w", err)
		return err
	}
	defer obj.Body.Close()

	outFilePath := filepath.Join(dir, filepath.Base(*item.Key))
	outFile, err := os.Create(outFilePath)
	if err != nil {
		logger.Errorf("dto creating file %s: %w", outFilePath, err)
		return err
	}
	defer outFile.Close()

	if _, err := io.Copy(outFile, obj.Body); err != nil {
		logger.Errorf("dto writing file content: %w", err)
		return err
	}
	logger.Infof("Successfully downloaded file: %s", outFilePath)
	return nil
}
