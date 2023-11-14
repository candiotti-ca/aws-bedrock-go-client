package awsbedrockgoclient

import (
	"aws-bedrock-go-client/models"

	"github.com/aws/aws-sdk-go-v2/aws"
)

type AnthropicInput struct {
	Prompt        string   `json:"prompt"`
	Temperature   float64  `json:"temperature"`
	TopP          float64  `json:"top_p"`
	TopK          float64  `json:"top_k"`
	MaxTokens     int      `json:"max_tokens_to_sample"`
	StopSequences []string `json:"stop_sequences"`
}

type AnthropicOutput struct {
	Completion   string `json:"completion"`
	FinishReason string `json:"stop_reason"`
	Model        string `json:"model"`
}

// https://docs.anthropic.com/claude/reference/complete_post
func NewAnthropic(cfg aws.Config, model models.Model) Client[AnthropicInput, AnthropicOutput] {
	return New[AnthropicInput, AnthropicOutput](cfg, model)
}
