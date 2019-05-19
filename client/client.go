package client

// untested, still...

import (
	"github.com/aaronland/go-artisanal-integers"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"strconv"
)

type LambdaClient struct {
	artisanalinteger.Client
	service     *lambda.Lambda
	lambda_func string
}

type Integer struct {
	Integer int64 `json:"integer"`
}

func NewLambdaClient(sess *session.Session, lambda_func string) (artisanalinteger.Client, error) {

	service := lambda.New(sess)

	cl := LambdaClient{
		service:     service,
		lambda_func: lambda_func,
	}

	return &cl, nil
}

func (cl *LambdaClient) NextInt() (int64, error) {

	input := &lambda.InvokeInput{
		FunctionName:   aws.String(cl.lambda_func),
		InvocationType: aws.String("RequestResponse"),
	}

	rsp, err := cl.service.Invoke(input)

	if err != nil {
		return -1, err
	}

	str_i := string(rsp.Payload)
	i, err := strconv.ParseInt(str_i, 10, 64)

	return i, nil
}
