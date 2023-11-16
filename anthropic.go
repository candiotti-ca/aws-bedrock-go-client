package awsbedrockgoclient

import (
	"github.com/aws/aws-sdk-go-v2/aws"
)

// Complete documentation on https://docs.anthropic.com/claude/reference/complete_post
type AnthropicInput struct {
	Prompt string `json:"prompt"`
	// Use a lower temperature to decrease randomness in the response
	Temperature float64 `json:"temperature" validate:"min=0,max=1"`
	// Use a lower value to ignore less probable responses
	TopP float64 `json:"top_p" validate:"min=0,max=1"`
	// Specify the number of token choices the model uses to generate the next token
	TopK      float64 `json:"top_k" validate:"min=0,max=500"`
	MaxTokens int     `json:"max_tokens_to_sample" validate:"min=0,max=4096"`
	// Character sequences that stops generation. Use pipe to separate sequences
	StopSequences []string `json:"stop_sequences"`
}

type AnthropicOutput struct {
	Completion   string `json:"completion"`
	FinishReason string `json:"stop_reason"`
	Model        string `json:"model"`
}

func NewAnthropicClaudeInstantV1(cfg aws.Config) Client[AnthropicInput, AnthropicOutput] {
	return New[AnthropicInput, AnthropicOutput](cfg, anthropicClaudeInstantV1)
}

func NewAnthropicClaudeV1(cfg aws.Config) Client[AnthropicInput, AnthropicOutput] {
	return New[AnthropicInput, AnthropicOutput](cfg, anthropicClaudeV1)
}

func NewAnthropicClaudeV2(cfg aws.Config) Client[AnthropicInput, AnthropicOutput] {
	return New[AnthropicInput, AnthropicOutput](cfg, anthropicClaudeV2)
}
