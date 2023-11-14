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
	bedrockClient     *bedrockruntime.Client
	model             string
	supportsStreaming bool
}

// https://docs.aws.amazon.com/bedrock/latest/APIReference
func New[I, O any](cfg aws.Config, model models.Model) Client[I, O] {
	b := bedrockruntime.NewFromConfig(cfg)
	return Client[I, O]{
		bedrockClient:     b,
		model:             string(model),
		supportsStreaming: doesSupportStreaming(model),
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

func (c Client[I, O]) QueryStream(ctx context.Context, input I) (*bedrockruntime.InvokeModelWithResponseStreamOutput, error) {
	if !c.supportsStreaming {
		return nil, fmt.Errorf("model <%s> doesn't support streaming", string(c.model))
	}

	body, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal input: %w", err)
	}

	output, err := c.bedrockClient.InvokeModelWithResponseStream(ctx, &bedrockruntime.InvokeModelWithResponseStreamInput{
		ModelId:     &c.model,
		Body:        body,
		Accept:      aws.String("*/*"),
		ContentType: aws.String("application/json"),
	})
	if err != nil {
		return nil, fmt.Errorf("invoke model request has failed: %w", err)
	}

	return output, nil
}

func doesSupportStreaming(model models.Model) bool {
	return model == models.Anthropic_ClaudeInstantV1 ||
		model == models.Anthropic_ClaudeV1 ||
		model == models.Anthropic_ClaudeV2 ||
		model == models.Meta_Llama2ChatV1 ||
		model == models.Amazon_TitanTextLarge ||
		model == models.Amazon_TitanTextLiteV1 ||
		model == models.Amazon_TitanTextExpressV1 ||
		model == models.Amazon_TitanTextEmbeddingsV1 ||
		model == models.Cohere_CommandV14
}
