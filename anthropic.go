package awsbedrockgoclient

// Complete documentation on https://docs.anthropic.com/claude/reference/complete_post
type AnthropicInput struct {
	Prompt string `json:"prompt"`
	// Use a lower temperature to decrease randomness in the response
	Temperature float64 `json:"temperature" validate:"min=0,max=1" default:"0.5"`
	// Use a lower value to ignore less probable responses
	TopP float64 `json:"top_p" validate:"min=0,max=1" default:"1"`
	// Specify the number of token choices the model uses to generate the next token
	TopK      float64 `json:"top_k" validate:"min=0,max=500" default:"250"`
	MaxTokens int     `json:"max_tokens_to_sample" validate:"min=0,max=4096" default:"200"`
	// Character sequences that stops generation. Use pipe to separate sequences
	StopSequences []string `json:"stop_sequences"`
}

type AnthropicOutput struct {
	Completion   string `json:"completion"`
	FinishReason string `json:"stop_reason"`
	Model        string `json:"model"`
}

func NewAnthropicClaudeInstantV1(b BedrockRuntime) Client[AnthropicInput, AnthropicOutput] {
	return New[AnthropicInput, AnthropicOutput](b, anthropicClaudeInstantV1)
}

func NewAnthropicClaudeV1(b BedrockRuntime) Client[AnthropicInput, AnthropicOutput] {
	return New[AnthropicInput, AnthropicOutput](b, anthropicClaudeV1)
}

func NewAnthropicClaudeV2(b BedrockRuntime) Client[AnthropicInput, AnthropicOutput] {
	return New[AnthropicInput, AnthropicOutput](b, anthropicClaudeV2)
}
