package handlers

import (
	"fmt"
	"net/http"

	"github.com/daveod/aws-lambda-in-go-lang/pkg/team"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var ErrorMethodNotAllowed = "method Not allowed"

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

func GetTeam(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse,
	error,
) {
	fmt.Println("In GetTeam")
	nickName := req.QueryStringParameters["nickName"]
	fmt.Printf("nickName query = %s\n", nickName)
	if len(nickName) > 0 {
		// Get single Team
		result, err := team.FetchTeam(nickName, tableName, dynaClient)
		if err != nil {
			return err.Error()
		}

		return result
	}

	// Get list of Teams
	result, err := team.FetchTeams(tableName, dynaClient)
	if err != nil {
		return err.Error()
	}

	return result
}

func CreateTeam(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse,
	error,
) {
	result, err := team.CreateTeam(req, tableName, dynaClient)
	if err != nil {
		return err.Error()
	}
	return result
}

func UpdateTeam(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse,
	error,
) {
	result, err := team.UpdateTeam(req, tableName, dynaClient)
	if err != nil {
		return err.Error()
	}
	return result
}

func DeleteTeam(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse,
	error,
) {
	err := team.DeleteTeam(req, tableName, dynaClient)
	if err != nil {
		return err.Error()
	}
	return nil
}

func UnhandledMethod() (*events.APIGatewayProxyResponse, error) {
	return ErrorMethodNotAllowed
}
