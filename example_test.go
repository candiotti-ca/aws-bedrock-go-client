package awsbedrockgoclient_test

import (
	"context"
	"fmt"

	awsbedrockgoclient "github.com/candiotti-ca/aws-bedrock-go-client"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

func ExampleQuery() {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic(err)
	}

	// COHERE COMMAND MODEL
	client := awsbedrockgoclient.NewCohereCommandV14(bedrockruntime.NewFromConfig(cfg))

	resp, err := client.Query(ctx, awsbedrockgoclient.CohereCommandInput{
		Prompt: "Hello",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("[COHERE] resp: %v\n", resp)
}
