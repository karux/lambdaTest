package main

//Author: Chris Haddad   haddadc@karux.net
//Copyright (c) Karux LLC 2018

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/lambdacontext"

	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Message string `json:"message"`
}

type Request map[string]string

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
	if rq != nil {
		log.Println(rq)
		clientId = rq["clientid"]
		requestMapJSON, err = json.Marshal(rq)
	}

	return Response{
		Message: "{ handler: 'world', clientId: '" + clientId + "', context: " + string(contextJSON) + ", requestMap: " + string(requestMapJSON) + ", deadline: '" + time.Until(deadline).String() + "'  }",
		// echo request parameters back
	}, nil
}

func main() {
	lambda.Start(Handler)
}
