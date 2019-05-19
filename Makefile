deps:
	github.com/aaronland/go-artisanal-integers v0.1.0
	github.com/whosonfirst/algnhsa v0.1.0
	github.com/aws/aws-sdk-go/aws v1.19.33

fmt:
	go fmt client/*.go
	go fmt cmd/*.go
	go fmt server/*.go

tools: 
	if test ! -d bin; then mkdir bin; fi
	go build -o bin/int cmd/int/main.go
