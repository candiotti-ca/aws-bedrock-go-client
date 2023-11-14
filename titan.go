package awsbedrockgoclient

import (
	"aws-bedrock-go-client/models"

	"github.com/aws/aws-sdk-go-v2/aws"
)

type TitanInput struct {
	Prompt string `json:"inputText"`
	Config struct {
		Temperature   float64  `json:"temperature"`
		TopP          float64  `json:"topP"`
		MaxTokenCount int      `json:"maxTokenCount"`
		StopSequences []string `json:"stopSequences"`
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

func NewTitan(cfg aws.Config, model models.Model) Client[TitanInput, TitanOutput] {
	return New[TitanInput, TitanOutput](cfg, model)
}
