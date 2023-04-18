package opens3

import (
	"os"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	_, err := New("us-west-2")
	if err != nil {
		t.Errorf("New() failed: %s", err)
	}
}

func getS3goClient(t *testing.T) *S3go {
	client, err := New("us-west-2")
	if err != nil {
		t.Fatalf("Failed to create S3go client: %s", err)
	}
	return client
}

func TestListBuckets(t *testing.T) {
	client := getS3goClient(t)
	_, err := client.ListBuckets()
	if err != nil {
		t.Errorf("ListBuckets() failed: %s", err)
	}
}

func TestCreateAndDeleteBucket(t *testing.T) {
	client := getS3goClient(t)
	bucketName := "go-s3go-test-bucket"

	// Test CreateBucket
	_, err := client.CreateBucket(bucketName)
	if err != nil {
		t.Errorf("CreateBucket() failed: %s", err)
	}

	// Verify that the bucket was created
	buckets, err := client.ListBuckets()
	if err != nil {
		t.Fatalf("ListBuckets() failed: %s", err)
	}

	found := false
	for _, bucket := range buckets {
		if *bucket.Name == bucketName {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("CreateBucket() failed: bucket not found in the list of buckets")
	}

	// Test DeleteBucket
	_, err = client.DeleteBucket(bucketName)
	if err != nil {
		t.Errorf("DeleteBucket() failed: %s", err)
	}
}

