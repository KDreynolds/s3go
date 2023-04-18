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

func TestObjectOperations(t *testing.T) {
	client := getS3goClient(t)
	bucketName := "go-s3go-test-object-operations"
	testFileName := "testfile.txt"
	testFileKey := "testfile_key.txt"
	downloadedFileName := "downloaded_testfile.txt"

	// Create a test bucket
	_, err := client.CreateBucket(bucketName)
	if err != nil {
		t.Fatalf("CreateBucket() failed: %s", err)
	}
	defer client.DeleteBucket(bucketName) // Clean up the test bucket

	// Create a test file
	err = os.WriteFile(testFileName, []byte("This is a test file"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %s", err)
	}
	defer os.Remove(testFileName) // Clean up the test file

	// Test UploadFile
	_, err = client.UploadFile(bucketName, testFileName, testFileKey)
	if err != nil {
		t.Errorf("UploadFile() failed: %s", err)
	}

	// Test ListObjects
	objects, err := client.ListObjects(bucketName)
	if err != nil {
		t.Errorf("ListObjects() failed: %s", err)
	}

	found := false
	for _, object := range objects {
		if *object.Key == testFileKey {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("UploadFile() failed: object not found in the list of objects")
	}

	// Test DownloadFile
	_, err = client.DownloadFile(bucketName, testFileKey, downloadedFileName)
	if err != nil {
		t.Errorf("DownloadFile() failed: %s", err)
	}
	defer os.Remove(downloadedFileName) // Clean up the downloaded test file

	// Test DeleteObject
	_, err = client.DeleteObject(bucketName, testFileKey)
	if err != nil {
		t.Errorf("DeleteObject() failed: %s", err)
	}
}

func TestCopyObject(t *testing.T) {
	client := getS3goClient(t)
	bucketName := "go-s3go-test-copy-object"
	sourceKey := "source_key.txt"
	destinationKey := "destination_key.txt"

	// Create a test bucket
	_, err := client.CreateBucket(bucketName)
	if err != nil {
		t.Fatalf("CreateBucket() failed: %s", err)
	}
	defer client.DeleteBucket(bucketName) // Clean up the test bucket

	// Create a test file
	sourceFileName := "sourcefile.txt"
	err = os.WriteFile(sourceFileName, []byte("This is the source file"), 0644)
	if err != nil {
		t.Fatalf("Failed to create source file: %s", err)
	}
	defer os.Remove(sourceFileName) // Clean up the source test file

	// Upload the source file
	_, err = client.UploadFile(bucketName, sourceFileName, sourceKey)
	if err != nil {
		t.Fatalf("UploadFile() failed: %s", err)
	}

	// Test CopyObject
	_, err = client.CopyObject(bucketName, sourceKey, bucketName, destinationKey)
	if err != nil {
		t.Errorf("CopyObject() failed: %s", err)
	}

	// Check if the destination object exists
	objects, err := client.ListObjects(bucketName)
	if err != nil {
		t.Errorf("ListObjects() failed: %s", err)
	}
	
	found := false
	for _, object := range objects {
		if *object.Key == destinationKey {
			found = true
			break
		}
	}
	
	if !found {
		t.Errorf("CopyObject() failed: destination object not found in the list of objects")
	}
	
	// Clean up the test objects
	_, _ = client.DeleteObject(bucketName, sourceKey)
	_, _ = client.DeleteObject(bucketName, destinationKey)
	
