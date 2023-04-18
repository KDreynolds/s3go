package opens3

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
    "github.com/aws/aws-sdk-go/service/s3/s3manager"
    "github.com/aws/aws-sdk-go/service/s3/s3manager"
    "os"
	"time"
)

type S3go struct {
    svc *s3.S3
}

func New(region string) (*S3go, error) {
    sess, err := session.NewSession(&aws.Config{
	Region: aws.String(region),
    })

    if err != nil {
	return nil, err
    }

    svc := s3.New(sess)
    return &OpenS3{svc: svc}, nil
}

func (o *S3go) ListBuckets() ([]*s3.Bucket, error) {
	result, err := o.svc.ListBuckets(nil)
	if err != nil {
		return nil, err
	}

	return result.Buckets, nil
}

func (o *S3go) CreateBucket(bucketName string) (*s3.CreateBucketOutput, error) {
	input := &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	}
	result, err := o.svc.CreateBucket(input)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (o *S3go) DeleteBucket(bucketName string) (*s3.DeleteBucketOutput, error) {
	input := &s3.DeleteBucketInput{
		Bucket: aws.String(bucketName),
	}
	result, err := o.svc.DeleteBucket(input)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (o *S3go) ListObjects(bucketName string) ([]*s3.Object, error) {
    input := &s3.ListObjectsInput{
        Bucket: aws.String(bucketName),
    }
    result, err := o.svc.ListObjects(input)
    if err != nil {
        return nil, err
    }

    return result.Contents, nil
}

func (o *S3go) UploadFile(bucketName, filePath, key string) (*s3manager.UploadOutput, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	uploader := s3manager.NewUploader(o.svc.Session)
	input := &s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   file,
	}
	result, err := uploader.Upload(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			return nil, aerr
		}
		return nil, err
	}

	return result, nil
}

func (o *S3go) DownloadFile(bucketName, key, filePath string) (int64, error) {
	file, err := os.Create(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	downloader := s3manager.NewDownloader(o.svc.Session)
	input := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	}
	bytesWritten, err := downloader.Download(file, input)
	if err != nil {
		return 0, err
	}

	return bytesWritten, nil
}

func (o *S3go) DeleteObject(bucketName, key string) (*s3.DeleteObjectOutput, error) {
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	}
	result, err := o.svc.DeleteObject(input)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (o *S3go) CopyObject(sourceBucket, sourceKey, destinationBucket, destinationKey string) (*s3.CopyObjectOutput, error) {
	source := fmt.Sprintf("%s/%s", sourceBucket, sourceKey)
	input := &s3.CopyObjectInput{
		Bucket:     aws.String(destinationBucket),
		CopySource: aws.String(url.PathEscape(source)),
		Key:        aws.String(destinationKey),
	}

	result, err := o.svc.CopyObject(input)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (o *S3go) GeneratePresignedURL(bucketName, key string, expiration time.Duration) (string, error) {
	request, _ := o.svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})
	url, err := request.Presign(expiration)
	if err != nil {
		return "", err
	}

	return url, nil
}



