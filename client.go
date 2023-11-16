package awsbedrockgoclient

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/go-playground/validator/v10"
)

type Client[I, O any] struct {
	bedrockClient     *bedrockruntime.Client
	model             string
	supportsStreaming bool
	validate          *validator.Validate
}

// https://docs.aws.amazon.com/bedrock/latest/APIReference
func New[I, O any](cfg aws.Config, model Model) Client[I, O] {
	b := bedrockruntime.NewFromConfig(cfg)
	return Client[I, O]{
		bedrockClient:     b,
		model:             string(model),
		supportsStreaming: doesSupportStreaming(model),
		validate:          validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (c Client[I, O]) Query(ctx context.Context, input I) (*O, error) {
	if err := c.validate.Struct(input); err != nil {
		return nil, fmt.Errorf("invalid input: %w", err)
	}

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
	if err := c.validate.Struct(input); err != nil {
		return nil, fmt.Errorf("invalid input: %w", err)
	}

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

func doesSupportStreaming(model Model) bool {
	return model == anthropicClaudeInstantV1 ||
		model == anthropicClaudeV1 ||
		model == anthropicClaudeV2 ||
		model == metaLlama2ChatV1 ||
		model == amazonTitanTextLarge ||
		model == amazonTitanTextLiteV1 ||
		model == amazonTitanTextExpressV1 ||
		model == amazonTitanTextEmbeddingsV1 ||
		model == cohereCommandV14
}
