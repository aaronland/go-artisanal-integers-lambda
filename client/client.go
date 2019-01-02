package client

import (
	"github.com/aaronland/go-artisanal-integers"
)

type LambdaClient struct {
	artisanalinteger.Client
}

func NewLambdaClient() (artisanalinteger.Client, error) {

     cl := LambdaClient{

     }

     return &cl, nil
}

func (cl *LambdaClient) NextInt() (int64, error) {

}
