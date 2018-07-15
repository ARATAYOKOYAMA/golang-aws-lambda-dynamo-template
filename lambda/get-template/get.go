package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Request struct {
	ID    string `json:"id"`
}

type Response struct {
	Message string `json:"message"`
	Result Item `json:"item"`
	Ok      bool   `json:"ok"`
}

type Item struct {
	ID string`json:"id"`
	Name string`json:"name"`
}

//公式Doc https://docs.aws.amazon.com/ja_jp/sdk-for-go/v1/developer-guide/dynamo-example-read-table-item.html
func GetItem(request Request) (Response, error) {
	// session
	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	svc := dynamodb.New(sess)

	// GetItem
	//getParams := &dynamodb.GetItemInput{
	//	TableName: aws.String("go-demo"),
	//	Key: map[string]*dynamodb.AttributeValue{
	//		"id": {
	//			S: aws.String("1"),
	//		},
	//	},
	//}

	// GetItem parameter
	getParams := &dynamodb.GetItemInput{
		TableName: aws.String("go-demo"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(request.ID),
			},
		},
	}

	// エラー周り
	getItem, getErr := svc.GetItem(getParams)
	if getErr != nil {
		panic(getErr)
	}
	fmt.Println(getItem)

	// インスタンス立てた？
	item := Item{}

	// 紐づけた？
	err = dynamodbattribute.UnmarshalMap(getItem.Item, &item)

	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}


	return Response{
		Message: fmt.Sprintln(getItem.Item),
		Result: item,
		Ok:      true,
	}, nil
}

func main() {
	lambda.Start(GetItem)
}
