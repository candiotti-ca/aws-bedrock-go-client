package awsbedrockgoclient

// See full documentation on https://platform.stability.ai/docs/api-reference
type StabilityInput struct {
	Prompts []string `json:"text_prompts"`
	// Determines how much the final image portrays the prompt
	Scale float64 `json:"cfg_scale" validate:"min=0,max=30"`
	// Generation step determines how many times the image is sampled
	Steps int `json:"steps" validate:"min=10,max=150"`
	// The seed determines the initial noise setting
	Seed int `json:"seed"`
}

type StabilityOutput struct {
	Generations []struct {
		ImageB64     string `json:"base64"`
		FinishReason string `json:"finishReason"`
		Seed         int    `json:"seed"`
	} `json:"generations"`
}

func NewStableDiffusionXlV0(b BedrockRuntime) Client[StabilityInput, StabilityOutput] {
	return New[StabilityInput, StabilityOutput](b, stabilityAIStableDiffusionXLV0)
}

func NewStableDiffusionXlV1(b BedrockRuntime) Client[StabilityInput, StabilityOutput] {
	return New[StabilityInput, StabilityOutput](b, stabilityAIStableDiffusionXLV1)
}
