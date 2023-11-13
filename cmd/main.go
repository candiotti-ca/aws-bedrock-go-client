package main

import (
	awsbedrockgoclient "aws-bedrock-go-client"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithSharedConfigProfile("edfx"), config.WithRegion("us-east-1"))
	if err != nil {
		panic(err)
	}

	c := awsbedrockgoclient.New(cfg, awsbedrockgoclient.ModelCohereCommand)

	resp, err := c.Query("Hello")
	if err != nil {
		panic(err)
	}

	fmt.Printf("resp: %v\n", resp)
}
