// s3 API test.go file, for learning purpose.
package main

import (
	"fmt"
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
)

func main() {

	// Set up s3 connection
	// For security reason I have left the keys blank as default.
	auth := aws.Auth{
		AccessKey: "",
		SecretKey: "",
	}
	apnortheast := aws.USEast
	s3conn := s3.New(auth, apnortheast)

	// connection.Bucket into existing file.
	fmt.Println("Awesome Eric")
	mybucket := s3conn.Bucket("humancool")

	response, err := mybucket.List("", "", "", 1000)
	if err != nil {
		panic(err)
	}
	for _, v := range response.Contents {
		fmt.Println(v.Key)
	}

}
