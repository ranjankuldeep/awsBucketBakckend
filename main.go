package main

import (
	"fmt"
	"os"
)

// An error message to be displayed to standard output
func exitError(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, msg+"\n", args...)
	os.Exit(1)
}

func main() {
	// bucketName := "testinguploadkd"
	// awsbucket.CreateBucket(bucketName, "us-east-1")
	// awsbucket.ListBuckets()
	// awsbucket.UploadItem(bucketName, "randomfile.txt", "us-west-2")
}
