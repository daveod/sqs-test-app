package handlers

import (
	"errors"
	"fmt"

	"github.com/daveod/sqs-test-app/pkg/team"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var ErrorMethodNotAllowed = "method Not allowed"

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

func GetTeam(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*team.Team, error) {
	fmt.Println("In GetTeam")
	nickName := req.QueryStringParameters["nickName"]
	fmt.Printf("nickName query = %s\n", nickName)
	if len(nickName) > 0 {
		// Get single Team
		result, err := team.FetchTeam(nickName, tableName, dynaClient)
		if err != nil {
			return nil, err
		}

		return result, nil
	}

	// Get list of Teams
	result, err := team.FetchTeams(tableName, dynaClient)
	if err != nil {
		return nil, err
	}

	fmt.Printf("nGetTeam Result = %v\n", result)
	return nil, nil
}

func CreateTeam(body string, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*team.Team, error) {
	result, err := team.CreateTeam(body, tableName, dynaClient)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateTeam(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*team.Team, error) {
	result, err := team.UpdateTeam(req, tableName, dynaClient)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func DeleteTeam(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse,
	error,
) {
	err := team.DeleteTeam(req, tableName, dynaClient)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func UnhandledMethod() error {
	return errors.New(ErrorMethodNotAllowed)
}
