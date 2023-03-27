package main

import (
	"PMS/handlers"
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

//ginLambda is a pointer to a GinLambda object, which is a type provided by the ginadapter package
// router is a pointer to a gin.Engine object, which is the main router for the Gin web server
var ginLambda *ginadapter.GinLambda
var router *gin.Engine

func main() {
	const functionName = "main.main"
	fmt.Println(functionName)
	// it initializes the router
	router = gin.Default()
	//function that sets up the routes for the Gin server by http handlers for fun/methods.
	handlers.SetupRoutes(router)
	//This tells AWS to start the function and pass requests to the LambdaHandler() function.
	lambda.Start(LambdaHandler)
}

func LambdaHandler(ctx context.Context, input map[string]interface{}) (interface{}, error) {
	fmt.Println("inside LambdaHandler")
	// Convert the input to a JSON-encoded byte
	inputBytes, _ := json.Marshal(input)
	//variable to hold the API Gateway event
	var awsAPIGatewayEvent events.APIGatewayProxyRequest
	// storing the input into the awsAPIGatewayEvent
	err := json.Unmarshal(inputBytes, &awsAPIGatewayEvent)
	//if err is nil(no error) and httpmethod=string then input is http request from apigateway
	// stores the http request in awsEvent
	//create a new ginadpter with router object
	// func of ginlambda(proxyWithContext) is called with 2 objects init as arguments

	if err == nil && awsAPIGatewayEvent.HTTPMethod != "" {
		awsEvent := awsAPIGatewayEvent
		ginLambda = ginadapter.New(router)
		return ginLambda.ProxyWithContext(ctx, awsEvent)
	}

	return "invalid request type", nil
}
