// s3 API test.go file, for learning purpose.
package main

import (
	"fmt"
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
	"os"
)

func main() {

	// Set up s3 connection
	// For security reason I have left the keys blank as default.
	auth := aws.Auth{
		AccessKey: os.Getenv("AWS_ACCESS_KEY_ID"),
		SecretKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
	}
	USEast := aws.USEast
	s3conn := s3.New(auth, USEast)

	// connection.Bucket into existing file.
	fmt.Println("Awesome Eric")

	mybucket := s3conn.Bucket("humancool")

	mybucket.PutBucket("New Bucket")

	response, err := mybucket.List("magazines", "", "", 1000)
	if err != nil {
		panic(err)
	}
	for _, v := range response.Contents {
		fmt.Println(v.Key)
	}

}
