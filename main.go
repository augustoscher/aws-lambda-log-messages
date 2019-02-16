package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"bitbucket.org/augustoscher/lambda-log-erros/model"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleGetMensagemErro(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Retrieves the ID from the URL
	id := req.PathParameters["id"]

	// Fetches the requested Todo
	msg, err := GetMensagemErro(id)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       http.StatusText(http.StatusInternalServerError),
		}, nil
	}

	// Marshals the struct so the API Gateway is able to proccess it
	js, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       http.StatusText(http.StatusInternalServerError),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(js),
	}, nil
}

func handleGetMensagensErro(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	msgs, err := GetMensagensErro()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       http.StatusText(http.StatusInternalServerError),
		}, nil
	}

	js, err := json.Marshal(msgs)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       http.StatusText(http.StatusInternalServerError),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(js),
	}, nil
}

func handleCreateMensagemErro(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var msg model.MensagemErro
	err := json.Unmarshal([]byte(req.Body), &msg)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	err = CreateMensagemErro(msg)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       "Created",
	}, nil
}

func router(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	if req.Path == "/errors" {
		if req.HTTPMethod == "GET" {
			hasID, _ := regexp.MatchString("/errors/.+", req.Path)
			if hasID {
				return handleGetMensagemErro(req)
			}

			if req.Path == "/errors" {
				return handleGetMensagensErro(req)
			}
		}

		if req.HTTPMethod == "POST" {
			return handleCreateMensagemErro(req)
		}
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusMethodNotAllowed,
		Body:       http.StatusText(http.StatusMethodNotAllowed),
	}, nil
}

func main() {
	lambda.Start(router)
}
