package main

import (
	helper "github.com/dangula/goDockerTest/helpers"
	"fmt"
	"os"
)


// Lists all objects in a bucket using pagination
//
// Usage:
// listObjects <bucket>
func main() {

	fmt.Println("AWS_HOST:", os.Getenv("AWS_HOST"))
	fmt.Println("AWS_KEY:", os.Getenv("AWS_KEY"))
	fmt.Println("AWS_SECRET:", os.Getenv("AWS_SECRET"))



	created,err := helper.CreateBucket("myBucket1")
	if err !=nil{
		fmt.Println("error creating bucket")
		panic(err)
	}
	if !created {
		fmt.Printf("Object not created successfully")
	} else {
		fmt.Println("myBucket1 created successfully")
	}

	bucketPresent,err := helper.IsBucketPresent("myBucket1")
	if err !=nil{
		fmt.Println("error getting bucket")
		panic(err)
	}
	if !bucketPresent {
		fmt.Printf("Bucket not present")
	} else {
		fmt.Println("Bucket present")
	}

	putSuccess,err := helper.PutObjectInBucket("myBucket1","myobj1","object 1 datat for rook")
	if err !=nil{
		fmt.Println("error getting bucket")
		panic(err)
	}
	if !putSuccess {
		fmt.Printf("cannot put data in object store")
	} else {
		fmt.Println("put data in object store successfully")
	}

	data,err :=helper.GetObjectFromBucket("myBucket1","myobj1")
	if err !=nil{
		fmt.Println("error getting bucket")
		panic(err)
	}
	fmt.Printf("data from object store : ", data)

}