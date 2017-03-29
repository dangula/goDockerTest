package main

import (
	"log"
	"github.com/minio/minio-go"

)


// Lists all objects in a bucket using pagination
//
// Usage:
// listObjects <bucket>
func main() {
	// Note: YOUR-ACCESSKEYID, YOUR-SECRETACCESSKEY and my-bucketname are
	// dummy values, please replace them with original values.

	// Requests are always secure (HTTPS) by default. Set secure=false to enable insecure (HTTP) access.
	// This boolean value is the last argument for New().

	// New returns an Amazon S3 compatible client object. API compatibility (v2 or v4) is automatically
	// determined based on the Endpoint value.
	s3Client, err := minio.New("10.101.76.217:53390", "G4ZU50JKHOBG3N8NUH88",
		"gmNoXmRsJRJwwJ0KGCy4NCbbfOc7aR1sUSeUfVOu", false)
	if err != nil {
		log.Fatalln(err)
	}

	err = s3Client.MakeBucket("my-bucketname", "us-east-1")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Success")
}