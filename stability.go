package awsbedrockgoclient

import (
	"aws-bedrock-go-client/models"

	"github.com/aws/aws-sdk-go-v2/aws"
)

type StabilityInput struct {
	Prompts []string `json:"text_prompts"`
	Scale   float64  `json:"cfg_scale"`
	Steps   int      `json:"steps"`
	Seed    int      `json:"seed"`
}

type StabilityOutput struct {
	Generations []struct {
		ImageB64     string `json:"base64"`
		FinishReason string `json:"finishReason"`
		Seed         int    `json:"seed"`
	} `json:"generations"`
}

// https://platform.stability.ai/docs/api-reference
func NewStability(cfg aws.Config, model models.Model) Client[StabilityInput, StabilityOutput] {
	return New[StabilityInput, StabilityOutput](cfg, model)
}
