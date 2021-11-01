package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func DownloadObject(sess *session.Session) {

	file, err := os.Create("downloaded_data.txt")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer file.Close()

	downloader := s3manager.NewDownloader(sess)

	_, err = downloader.Download(file, &s3.GetObjectInput{
		Bucket: aws.String("supergolang"),
		Key: aws.String("myfile.txt"),
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Successfully downloaded")
}

func d () {

	fmt.Println("Download file")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1"),
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	DownloadObject(sess)
}
