package get_template

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	ID    int `json:"id"`
}

type Response struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

func GetItem() (Response, error) {
	// session
	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	svc := dynamodb.New(sess)

	// GetItem
	getParams := &dynamodb.GetItemInput{
		TableName: aws.String("go-demo"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String("1"),
			},
		},
	}

	getItem, getErr := svc.GetItem(getParams)
	if getErr != nil {
		panic(getErr)
	}
	fmt.Println(getItem)

	return Response{
		Message: fmt.Sprintln(getItem.Item),
		Ok:      true,
	}, nil
}

func main() {
	lambda.Start(GetItem)
}
