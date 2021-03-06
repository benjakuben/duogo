package main

import (
	"github.com/benjakuben/duogo/internal/app/graphql/schema"
	"github.com/benjakuben/duogo/internal/pkg/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

var graphql = new(handler.GraphQl)

func init() {
	graphql.BuildSchema(schema.Schema, &schema.QueryResolver{})
}

func main() {
	lambda.Start(graphql.Lambda)
}
