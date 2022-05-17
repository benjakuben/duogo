package main

import (
	"jakuben/duogo/internal/pkg/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

// func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
// 	ApiResponse := events.APIGatewayProxyResponse{}
// 	// Switch for identifying the HTTP request
// 	switch request.HTTPMethod {
// 	case "GET":
// 		// Obtain the QueryStringParameter
// 		name := request.QueryStringParameters["name"]
// 		if name != "" {
// 			ApiResponse = events.APIGatewayProxyResponse{Body: "¡Hola, " + name + ", bienvenido!", StatusCode: 200}
// 		} else {
// 			ApiResponse = events.APIGatewayProxyResponse{Body: "Error: Query Parameter name missing", StatusCode: 500}
// 		}

// 	case "POST":
// 		//validates json and returns error if not working
// 		err := fastjson.Validate(request.Body)

// 		if err != nil {
// 			body := "Error: Invalid JSON payload ||| " + fmt.Sprint(err) + " Body Obtained" + "||||" + request.Body
// 			ApiResponse = events.APIGatewayProxyResponse{Body: body, StatusCode: 500}
// 		} else {
// 			ApiResponse = events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}
// 		}

// 	}
// 	// Response
// 	return ApiResponse, nil
// }

var graphql = new(handler.GraphQl)

func main() {
	lambda.Start(graphql.Lambda)
}
