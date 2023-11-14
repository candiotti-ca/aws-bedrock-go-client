package awsbedrockgoclient

import (
	"aws-bedrock-go-client/models"
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

type Client[I, O any] struct {
	bedrockClient *bedrockruntime.Client
	model         string
}

func New[I, O any](cfg aws.Config, model models.Model) Client[I, O] {
	b := bedrockruntime.NewFromConfig(cfg)
	return Client[I, O]{
		bedrockClient: b,
		model:         string(model),
	}
}

func (c Client[I, O]) Query(ctx context.Context, input I) (*O, error) {
	body, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal input: %w", err)
	}

	output, err := c.bedrockClient.InvokeModel(ctx, &bedrockruntime.InvokeModelInput{
		ModelId:     &c.model,
		Body:        body,
		Accept:      aws.String("*/*"),
		ContentType: aws.String("application/json"),
	})
	if err != nil {
		return nil, fmt.Errorf("invoke model request has failed: %w", err)
	}

	var response *O
	err = json.Unmarshal(output.Body, response)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal response: %w", err)
	}

	return response, nil
}
