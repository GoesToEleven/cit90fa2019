package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(whatever)
}

func whatever() (string, error) {
	return fmt.Sprintln("Hello world"), nil
}
