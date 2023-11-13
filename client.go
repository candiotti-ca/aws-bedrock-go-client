package awsbedrockgoclient

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/google/uuid"
)

type Model string

const ModelCohereCommand Model = "cohere.command-text-v14"

type InvokeModelBody struct {
	Prompt string `json:"prompt"`
}

type InvokeModelResponse struct {
	ID          uuid.UUID `json:"id"`
	Prompt      string    `json:"prompt"`
	Generations []struct {
		FinishReason string    `json:"finish_reason"`
		ID           uuid.UUID `json:"id"`
		Text         string    `json:"text"`
	} `json:"generations"`
}

type Client struct {
	bedrockClient *bedrockruntime.Client
	model         string
}

func New(cfg aws.Config, model Model) Client {
	b := bedrockruntime.NewFromConfig(cfg)
	return Client{
		bedrockClient: b,
		model:         string(model),
	}
}

// Query allows to query Aws Bedrock service with given prompt.
func (c Client) Query(prompt string) (string, error) {
	body, err := json.Marshal(InvokeModelBody{Prompt: prompt})
	if err != nil {
		return "", err
	}

	output, err := c.bedrockClient.InvokeModel(context.Background(), &bedrockruntime.InvokeModelInput{
		ModelId:     &c.model,
		Body:        []byte(body),
		Accept:      aws.String("*/*"),
		ContentType: aws.String("application/json"),
	})

	if err != nil {
		return "", fmt.Errorf("invoke model request has failed: %w", err)
	}

	response := InvokeModelResponse{}
	err = json.Unmarshal(output.Body, &response)
	if err != nil {
		return "", err
	}

	return response.Generations[0].Text, nil
}
