package handlers

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"goServerlessYT/pkg/user"
	"net/http"
)

var ErrorMethodNotAllowed = "method not allowed"

type ErrorBody struct {
	ErrorMsg *string `bson:"error,omitempty"`
}

func GetUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	email := req.QueryStringParameters["email"]

	if len(email) > 0 {
		result, err := user.FetchUser(email, tableName, dynaClient)
		if err != nil {
			return apiResponce(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
		}
		return apiResponce(http.StatusOK, result)
	}

	result, err := user.FetchUsers(tableName, "", dynaClient)
	if err != nil {
		return apiResponce(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return apiResponce(http.StatusOK, result)
}

func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	result, err := user.CreateUser(req, tableName, dynaClient)
	if err != nil {
		return apiResponce(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return apiResponce(http.StatusCreated, result)
}

func UpdateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	result, err := user.UpdateUser(req, tableName, dynaClient)
	if err != nil {
		return apiResponce(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return apiResponce(http.StatusOK, result)
}

func DeleteUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	err := user.DeleteUser(req, tableName, dynaClient)
	if err != nil {
		return apiResponce(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return apiResponce(http.StatusOK, nil)
}

func UnhandledMethod() (*events.APIGatewayProxyResponse, error) {
	return apiResponce(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
}
