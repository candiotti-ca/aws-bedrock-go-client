package awsbedrockgoclient

import (
	"github.com/aws/aws-sdk-go-v2/aws"
)

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

func NewTitanEmbeddingV1(cfg aws.Config) Client[TitanInput, TitanOutput] {
	return New[TitanInput, TitanOutput](cfg, amazonTitanTextEmbeddingsV1)
}

func NewTitanEmbeddingV2(cfg aws.Config) Client[TitanInput, TitanOutput] {
	return New[TitanInput, TitanOutput](cfg, amazonTitanTextEmbeddingsV2)
}

func NewTitanExpressV1(cfg aws.Config) Client[TitanInput, TitanOutput] {
	return New[TitanInput, TitanOutput](cfg, amazonTitanTextExpressV1)
}

func NewTitanTextLarge(cfg aws.Config) Client[TitanInput, TitanOutput] {
	return New[TitanInput, TitanOutput](cfg, amazonTitanTextLarge)
}

func NewTitanTextLiteV1(cfg aws.Config) Client[TitanInput, TitanOutput] {
	return New[TitanInput, TitanOutput](cfg, amazonTitanTextLiteV1)
}
