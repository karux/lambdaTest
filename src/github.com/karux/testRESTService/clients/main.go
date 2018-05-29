package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

type User struct {
	Name   string
	Skills string
	Active bool
	Exp    int
}

type Response interface{}

func main() {
	var req1 *User

	req1 = &User{}
	req1.Name = "arh111"
	req1.Skills = "javascript"
	req1.Active = false
	req1.Exp = 506
	fmt.Println("user:", req1)
	// marshall the request object into json payload
	payload, err := json.Marshal(req1)
	if err != nil {
		log.Fatal("could not serialize request object")
	}
	fmt.Println("payloadString:", string(payload))

	var payload64 string
	payload64 = base64.StdEncoding.EncodeToString(payload)
	fmt.Println("payload64:", payload64)

	// Create Lambda service client
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	client := lambda.New(sess, &aws.Config{Region: aws.String("us-east-1")})

	result, err := client.Invoke(&lambda.InvokeInput{
		FunctionName: aws.String("testRESTService-dev-iface"),
		//		ClientContext: aws.String("SLSDriver"),
		LogType: aws.String("Tail"),
		Payload: []byte(payload)})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			fmt.Println("Error calling iface ", aerr.Error(), " code:", aerr.Code(), "  ", err)
		}

		os.Exit(0)
	} else {
		fmt.Println("aws function response: " + string(result.Payload))
	}

}
