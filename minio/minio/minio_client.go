package minio

import (
	"log"
	"github.com/minio/minio-go"
)

const (
	endpoint        = "localhost:9000"
	accessKeyID     = "AKIAIOSFODNN7EXAMPLE"
	secretAccessKey = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
	useSSL          = false
)

//contentType := "application/octet-stream"

func FPutObject(bucketName string, location string, objectName string, filePath string, contentType string) (objectPath string) {
	client, err := minIoClient()
	// Make a new bucket called mymusic.
	bucketName, _ = bucketExists(client, bucketName, location)
	log.Println("BucketName " + bucketName)

	// Upload the zip file with FPutObject
	fileSize, err := client.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, fileSize)
	return bucketName + "/" + objectName
}
func bucketExists(client *minio.Client, bucketName string, location string) (string, error) {
	err := client.MakeBucket(bucketName, location)
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := client.BucketExists(bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}
	return bucketName, err
}

func minIoClient() (*minio.Client, error) {
	client, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}
	return client, err
}