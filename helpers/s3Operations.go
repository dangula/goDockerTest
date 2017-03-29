package helpers

import (
	"github.com/minio/minio-go"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"os"
	"io"
)

 var AWS_HOST = os.Getenv("AWS_ENDPOINT")
 var ACCCESS_KEY_ID = os.Getenv("AWS_KEY")
 var ACCESS_KEY_SECRET= os.Getenv("AWS_SECRET")

func CreateBucket(bucketName string) (bool,error){
	s3Client, err := minio.New(AWS_HOST, ACCCESS_KEY_ID, ACCESS_KEY_SECRET, false)
	if err != nil {
		return false,err
	}

	err = s3Client.MakeBucket(bucketName,"us-east-1")
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
	OBJ_PUT_FILE_NAME := uuid.NewV4().String()
        WriteToFile("/tmp",OBJ_PUT_FILE_NAME,data);

	_, err = s3Client.FPutObject(bucketName, objectName, "/tmp/"+OBJ_PUT_FILE_NAME,"application/text")
	if err != nil {
		return false,err
	}

	return true,nil
}


func GetObjectFromBucket(bucketName string,objectName string) (string,error){
	OBJ_GET_FILE_NAME := uuid.NewV4().String()
	s3Client, err := minio.New(AWS_HOST, ACCCESS_KEY_ID, ACCESS_KEY_SECRET, false)
	if err != nil {
		panic(err)
	}
	reader, err := s3Client.GetObject(bucketName, objectName)
	if err != nil {
		panic(err)
	}
	defer reader.Close()
	localFile, err := os.Create("/tmp/"+OBJ_GET_FILE_NAME)
	if err != nil {
		panic(err)
	}
	defer localFile.Close()

	stat, err := reader.Stat()
	if err != nil {
		panic(err)
	}

	if _, err := io.CopyN(localFile, reader, stat.Size); err != nil {
		panic(err)
	}

	str, err := ioutil.ReadFile("/tmp/"+OBJ_GET_FILE_NAME)
	if err != nil {
		panic(err)
	}

	return string(str),nil


}


