package main

import (
	"fmt"
	"os"
	"io/ioutil"
	helper "github.com/dangula/goDockerTest/helpers"
)



func main() {

	fmt.Println("AWS_ENDPOINT:", os.Getenv("AWS_ENDPOINT"))
	fmt.Println("AWS_KEY:", os.Getenv("AWS_KEY"))
	fmt.Println("AWS_SECRET:", os.Getenv("AWS_SECRET"))

	created,err := helper.CreateBucket("mybucket1")
	if err !=nil{
		fmt.Println("error creating bucket")
		panic(err)
	}
	if !created {
		fmt.Printf("Object not created successfully")
	} else {
		fmt.Println("myBucket1 created successfully")
	}

	bucketPresent,err := helper.IsBucketPresent("mybucket1")
	if err !=nil{
		fmt.Println("error getting bucket")
		panic(err)
	}
	if !bucketPresent {
		fmt.Printf("Bucket not present")
	} else {
		fmt.Println("Bucket present")
	}

	putSuccess,err := helper.PutObjectInBucket("mybucket1","myobj1","object 1 datat for rook")
	if err !=nil{
		fmt.Println("error getting bucket")
		panic(err)
	}
	if !putSuccess {
		fmt.Printf("cannot put data in object store")
	} else {
		fmt.Println("put data in object store successfully")
	}

	data,err :=helper.GetObjectFromBucket("mybucket1","myobj1")
	if err !=nil{
		fmt.Println("error getting bucket")
		panic(err)
	}
	fmt.Printf("data from object store : ", data)
	fmt.Println("    ")


	str,err :=helper.GetObjectFromBucket("mybucket1","myobj1")


	fmt.Println("get data : ",str)
}