package main

import (
	"context"
	"encoding/json"
	"log"
	"reflect"
	"time"

	"github.com/aws/aws-lambda-go/lambdacontext"

	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Message string `json:"message"`
}

type Request interface{}

func Handler(ctx context.Context, rq Request) (Response, error) {
	var requestMapJSON []byte
	var contextJSON []byte
	var err error

	deadline, _ := ctx.Deadline()

	log.Println(ctx)
	lc, _ := lambdacontext.FromContext(ctx)
	log.Println(lc)

	contextJSON, err = json.Marshal(lc)
	if err != nil {
		log.Println("context did not marshall into json")
	}

	var clientId string
	var objectValue string
	var objectType string
	if rq != nil {
		log.Println(rq)
		clientId = "index" // rq["clientid"]
		requestMapJSON, err = json.Marshal(rq)
		objectValue = reflect.ValueOf(rq).String()
		objectType = reflect.TypeOf(rq).String()
		// reflect on rq and determine base struct, then conditionally switch
		log.Println("value:", objectValue)
		log.Println("type:", objectType)

	}

	return Response{
		Message: "{ handler: 'iface', objectValue:'" + objectValue +
			"', objectType:'" + objectType +
			"' , clientId: '" + clientId +
			"', context: " + string(contextJSON) +
			", requestMap: " + string(requestMapJSON) +
			", deadline: '" + time.Until(deadline).String() +
			"'  }",
		// echo request parameters back
	}, nil
}

func main() {
	lambda.Start(Handler)
}
