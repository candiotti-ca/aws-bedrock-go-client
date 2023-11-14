package awsbedrockgoclient

import (
	"aws-bedrock-go-client/models"

	"github.com/aws/aws-sdk-go-v2/aws"
)

type MetaInput struct {
	Prompt       string  `json:"prompt"`
	Temperature  float64 `json:"temperature"`
	TopP         float64 `json:"top_p"`
	MaxGenLength int     `json:"max_gen_len"`
}

type MetaOutput struct {
	Generation           string `json:"generation"`
	PromptTokenCount     int    `json:"prompt_token_count"`
	GenerationTokenCount int    `json:"generation_token_count"`
	StopReason           string `json:"stop_reason"`
}

func NewMeta(cfg aws.Config, model models.Model) Client[MetaInput, MetaOutput] {
	return New[MetaInput, MetaOutput](cfg, model)
}
