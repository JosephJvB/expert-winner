package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func init() {
	// do startup work here, init clients etc
}

// json struct tag is what json input property name will be
// Actually property name is how to reference in code
// imo make it similar (the same) or get confused
type MyEvent struct {
	Name string `json:"Name"`
	Age int `json:"Age"`
}
 
type MyResponse struct {
	Message string `json:"Message"`
}
 
func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	fmt.Println("--- event ---")
	fmt.Println(event)
	fmt.Println(event.Name, event.Age)
	return MyResponse {
		Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age),
	}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}