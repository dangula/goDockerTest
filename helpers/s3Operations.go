package helpers

import (
	"github.com/minio/minio-go"
	"io/ioutil"
	"os"
)

const AWS_HOST = "10.101.76.217:53390"
const ACCCESS_KEY_ID = "G4ZU50JKHOBG3N8NUH88"
const ACCESS_KEY_SECRET= "gmNoXmRsJRJwwJ0KGCy4NCbbfOc7aR1sUSeUfVOu"

func CreateBucket(bucketName string) (bool,error){
	s3Client, err := minio.New(AWS_HOST, ACCCESS_KEY_ID, ACCESS_KEY_SECRET, false)
	if err != nil {
		return false,err
	}

	err = s3Client.MakeBucket(bucketName,"")
	if err != nil {
		return false,err
	}
	return true,nil
}

func IsBucketPresent(bucketName string) (bool,error){
	s3Client, err := minio.New(AWS_HOST, ACCCESS_KEY_ID, ACCESS_KEY_SECRET, false)
	if err != nil {
		return false,err
	}
	found,err :=s3Client.BucketExists(bucketName)

	return found,err
}

func PutObjectInBucket(bucketName string,objectName string,data string) (bool,error){
	s3Client, err := minio.New(AWS_HOST, ACCCESS_KEY_ID, ACCESS_KEY_SECRET, false)
	if err != nil {
		return false,err
	}
	object, err := os.Open("my-testfile")
	if err != nil {
		return false,err
	}
	defer object.Close()
	object.WriteString(data)

	_, err = s3Client.PutObject(bucketName, objectName, object, "application/octet-stream")
	if err != nil {
		return false,err
	}

	return true,nil
}

func GetObjectFromBucket(bucketName string,objectName string) (string,error){
	s3Client, err := minio.New(AWS_HOST, ACCCESS_KEY_ID, ACCESS_KEY_SECRET, false)
	if err != nil {
		return "error connecting to host",err
	}
	if err := s3Client.FGetObject(bucketName, objectName,"/tmp/getObj"); err != nil {
		return "failed to get object",err
	}
	str, err := ioutil.ReadFile("/tmp/getObj")
	if err != nil {
		return "fail", err
	}

	return string(str), nil
}

