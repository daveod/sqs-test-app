package main

import (
	"context"
	"fmt"
	"os"

	"github.com/daveod/sqs-test-app/pkg/handlers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var (
	dynaClient dynamodbiface.DynamoDBAPI
)

const tableName = "TeamsTable"

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {

	for _, message := range sqsEvent.Records {
		fmt.Printf("The message %s for event source %s = %s \n", message.MessageId, message.EventSource, message.Body)

		handlers.CreateTeam(message.Body, tableName, dynaClient)

		/*switch message.Body {
		case "GET":
			return handlers.GetTeam(req, tableName, dynaClient)
		case "POST":
			return handlers.CreateTeam(req, tableName, dynaClient)
		case "PUT":
			return handlers.UpdateTeam(req, tableName, dynaClient)
		case "DELETE":
			return handlers.DeleteTeam(req, tableName, dynaClient)
		default:
			return handlers.UnhandledMethod()
		}*/
	}

	return nil
}

func main() {
	fmt.Println("We are in the main function")
	region := os.Getenv("AWS_REGION")
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		return
	}
	dynaClient = dynamodb.New(awsSession)

	lambda.Start(handler)
}
