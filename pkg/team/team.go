package Team

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/daveod/sqs-test-app/pkg/validators"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var (
	ErrorFailedToUnmarshalRecord = "failed to unmarshal record"
	ErrorFailedToFetchRecord     = "failed to fetch record"
	ErrorInvalidTeamData         = "invalid Team data"
	ErrorInvalidEmail            = "invalid email"
	ErrorCouldNotMarshalItem     = "could not marshal item"
	ErrorCouldNotDeleteItem      = "could not delete item"
	ErrorCouldNotDynamoPutItem   = "could not dynamo put item error"
	ErrorTeamAlreadyExists       = "Team.Team already exists"
	ErrorTeamDoesNotExists       = "Team.Team does not exist"
)

type Team struct {
	NickName     string `json:"nickName"`
	City string `json:"city"`
	ShortName  string `json:"shortName"`
}

func FetchTeam(nickName, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*Team, error) {
	fmt.Println("In FetchTeam")
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"nickName": {
				S: aws.String(nickName),
			},
		},
		TableName: aws.String(tableName),
	}

	result, err := dynaClient.GetItem(input)
	fmt.Printf("err = %v\n", err)
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)

	}

	item := new(Team)
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		return nil, errors.New(ErrorFailedToUnmarshalRecord)
	}
	fmt.Printf("result = %v\n", item)
	return item, nil
}

func FetchTeams(tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*[]Team, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}
	result, err := dynaClient.Scan(input)
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}
	item := new([]Team)
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, item)
	return item, nil
}

func CreateTeam(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*Team,
	error,
) {
	fmt.Println("In CreateTeam")
	fmt.Printf("req.Body = %v\n", req.Body)

	var t Team
	if err := json.Unmarshal([]byte(req.Body), &t); err != nil {
		fmt.Printf("err = %v\n", err)
		return nil, errors.New(ErrorInvalidTeamData)
	}
	if !validators.IsEmailValid(t.NickName) {
		return nil, errors.New(ErrorInvalidEmail)
	}
	// Check if Team exists
	currentTeam, _ := FetchTeam(t.NickName, tableName, dynaClient)
	if currentTeam != nil && len(currentTeam.NickName) != 0 {
		return nil, errors.New(ErrorTeamAlreadyExists)
	}
	// Save Team

	av, err := dynamodbattribute.MarshalMap(t)
	if err != nil {
		return nil, errors.New(ErrorCouldNotMarshalItem)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = dynaClient.PutItem(input)
	if err != nil {
		return nil, errors.New(ErrorCouldNotDynamoPutItem)
	}
	return &t, nil
}

func UpdateTeam(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*Team,
	error,
) {
	var t Team
	fmt.Println("In UpdateTeam")

	if err := json.Unmarshal([]byte(req.Body), &t); err != nil {
		fmt.Printf("req.Body = %v\n", req.Body)
		fmt.Printf("Team = %v\n", t)
		return nil, errors.New(ErrorInvalidEmail)
	}

	// Check if Team exists
	currentTeam, _ := FetchTeam(t.NickName, tableName, dynaClient)
	if currentTeam != nil && len(currentTeam.NickName) == 0 {
		return nil, errors.New(ErrorTeamDoesNotExists)
	}

	// Save Team
	av, err := dynamodbattribute.MarshalMap(t)
	if err != nil {
		return nil, errors.New(ErrorCouldNotMarshalItem)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = dynaClient.PutItem(input)
	if err != nil {
		return nil, errors.New(ErrorCouldNotDynamoPutItem)
	}
	return &t, nil
}

func DeleteTeam(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) error {
	nickName := req.QueryStringParameters["nickName"]
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(nickName),
			},
		},
		TableName: aws.String(tableName),
	}
	_, err := dynaClient.DeleteItem(input)
	if err != nil {
		return errors.New(ErrorCouldNotDeleteItem)
	}

	return nil
}
