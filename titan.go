package awsbedrockgoclient

type TitanInput struct {
	Prompt string `json:"inputText"`
	Config struct {
		// Influences randomness response. A high temperature will give high random responses.
		Temperature float64 `json:"temperature" validate:"min=0,max=1"`
		// A high value tends to ignore less probable responses.
		TopP          float64 `json:"topP" validate:"min=0,max=1"`
		MaxTokenCount int     `json:"maxTokenCount" validate:"min=0,max=8000"`
		// Character sequences that stops generation. Use pipe to separate sequences.
		StopSequences []string `json:"stopSequences" validate:"len=20"`
	} `json:"textGenerationConfig"`
}

type TitanOutput struct {
	InputTextTokenCount int `json:"inputTextTokenCount"`
	Generations         []struct {
		TokenCount       int    `json:"tokenCount"`
		OutputText       string `json:"outputText"`
		CompletionReason string `json:"completionReason"`
	} `json:"results"`
}

func NewTitanEmbeddingV1(b BedrockRuntime) Client[TitanInput, TitanOutput] {
	return New[TitanInput, TitanOutput](b, amazonTitanTextEmbeddingsV1)
}

func NewTitanEmbeddingV2(b BedrockRuntime) Client[TitanInput, TitanOutput] {
	return New[TitanInput, TitanOutput](b, amazonTitanTextEmbeddingsV2)
}

func NewTitanExpressV1(b BedrockRuntime) Client[TitanInput, TitanOutput] {
	return New[TitanInput, TitanOutput](b, amazonTitanTextExpressV1)
}

func NewTitanTextLarge(b BedrockRuntime) Client[TitanInput, TitanOutput] {
	return New[TitanInput, TitanOutput](b, amazonTitanTextLarge)
}

func NewTitanTextLiteV1(b BedrockRuntime) Client[TitanInput, TitanOutput] {
	return New[TitanInput, TitanOutput](b, amazonTitanTextLiteV1)
}
