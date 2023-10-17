package main

import (
	"fmt"
	"os"

	"github.com/ranjankuldeep/awsbucket"
)

// An error message to be displayed to standard output
func exitError(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, msg+"\n", args...)
	os.Exit(1)
}

func main() {
	awsbucket.ListBuckets()
}
