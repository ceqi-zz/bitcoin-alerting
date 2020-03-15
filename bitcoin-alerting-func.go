package main

import (
	"log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
        "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/sns"
)

type Response struct {
	StatusCode int `json:"statusCode"`
	Body       string `json:"body"`

}

func HandleRequest() (Response, error) {
	sess, err := session.NewSession()
	if err != nil {
		log.Print("NewSession error:", err)
		return Response{StatusCode: 502}, nil
	}

	client := sns.New(sess)
	input := &sns.PublishInput{
		Message: aws.String("Alert: Bitcoin price is below $4000."),
		TopicArn: aws.String("arn:aws:sns:eu-west-1:182454815779:bitcoin-alerting-topic"),
	}

	result, err := client.Publish(input)
	if err!= nil {
		log.Print("Publish error:", err)
		return Response{StatusCode: 502}, nil
	}

	log.Print(result)

	log.Print("alert received.")
	return Response{StatusCode: 200, Body: "alert received"}, nil
}

func main() {
        lambda.Start(HandleRequest)
}
