package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadObject(sess *session.Session) {

	file, err := os.Open("myfile.txt")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer file.Close()

	uploader := s3manager.NewUploader(sess)

	result, err := uploader.Upload(&s3manager.UploadInput{
		ACL: aws.String("public-read"),
		Bucket: aws.String("supergolang"),
		Key: aws.String("myfile.txt"),
		Body: file,
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Successfully uploaded")
	fmt.Printf("Upload Result: %+v\n", result)
}

func u () {

	fmt.Println("Upload file")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1"),
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	UploadObject(sess)
}
