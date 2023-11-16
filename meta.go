package awsbedrockgoclient

type MetaInput struct {
	Prompt string `json:"prompt"`
	// Use a lower value to decrease randomness in the response.
	Temperature float64 `json:"temperature" validate:"min=0,max=1"`
	// Use a lower value to ignore less probable options
	TopP         float64 `json:"top_p" validate:"min=0,max=1"`
	MaxGenLength int     `json:"max_gen_len" validate:"min=1,max=2048"`
}

type MetaOutput struct {
	Generation           string `json:"generation"`
	PromptTokenCount     int    `json:"prompt_token_count"`
	GenerationTokenCount int    `json:"generation_token_count"`
	StopReason           string `json:"stop_reason"`
}

func NewMetaLlamaChatV1(b BedrockRuntime) Client[MetaInput, MetaOutput] {
	return New[MetaInput, MetaOutput](b, metaLlama2ChatV1)
}
