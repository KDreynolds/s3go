# s3go

`s3go` is a simple and easy-to-use Go library for managing Amazon S3 buckets and objects. It is inspired by the Python library `openS3`. This library provides a convenient interface for interacting with AWS S3 using the AWS SDK for Go.

## Features

- List buckets
- Create and delete buckets
- List objects in a bucket
- Upload and download files
- Delete and copy objects
- Generate presigned URLs

## Installation

To install `s3go`, run the following command:

```bash
go get -u github.com/KDreynolds/s3go
```bash

Usage

Here's a simple example demonstrating how to use the s3go library:


package main

import (
	"fmt"
	"log"
	"time"

	"github.com/your_username/s3go"
)

func main() {
	client, err := s3go.New("us-west-2")
	if err != nil {
		log.Fatal(err)
	}

	// List buckets
	buckets, err := client.ListBuckets()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Buckets:", buckets)

	// Other operations (create/delete buckets, list/upload/download/delete/copy objects, etc.)
	// ...
}

Documentation

For detailed documentation of the functions provided by s3go, please refer to the source code comments and the AWS SDK for Go documentation.
Contributing

Contributions are welcome! If you'd like to contribute, please fork the repository and create a pull request with your changes.
License

This project is licensed under the MIT License. See the LICENSE file for more information.

