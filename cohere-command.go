package awsbedrockgoclient

import (
	"aws-bedrock-go-client/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/google/uuid"
)

type CohereCommandReturnLikelihoods string

const (
	CohereCommandReturnLikelihoods_GENERATION CohereCommandReturnLikelihoods = "GENERATION"
	CohereCommandReturnLikelihoods_ALL        CohereCommandReturnLikelihoods = "ALL"
	CohereCommandReturnLikelihoods_NONE       CohereCommandReturnLikelihoods = "NONE"
)

type CohereCommandTruncate string

const (
	CohereCommandTruncate_NONE  CohereCommandTruncate = "NONE"
	CohereCommandTruncate_START CohereCommandTruncate = "START"
	CohereCommandTruncate_END   CohereCommandTruncate = "END"
)

type CohereCommandInput struct {
	Prompt            string                         `json:"prompt"`
	Temperature       float64                        `json:"temperature"`
	TopP              float64                        `json:"p"`
	TopK              float64                        `json:"k"`
	MaxReponseLength  int                            `json:"max_tokens"`
	StopSequences     []string                       `json:"stop_sequences"`
	ReturnLikelihoods CohereCommandReturnLikelihoods `json:"return_likelihoods"`
	Stream            bool                           `json:"stream"`
	NumGenerations    int                            `json:"num_generations"`
	LogitBias         map[string]float64             `json:"logit_bias"`
	Truncate          CohereCommandTruncate          `json:"truncate"`
}

type CohereCommandOutput struct {
	ID          uuid.UUID `json:"id"`
	Prompt      string    `json:"prompt"`
	Generations []struct {
		FinishReason string    `json:"finish_reason"`
		ID           uuid.UUID `json:"id"`
		Text         string    `json:"text"`
	} `json:"generations"`
}

func NewCohereCommandV14(cfg aws.Config) Client[CohereCommandInput, CohereCommandOutput] {
	return New[CohereCommandInput, CohereCommandOutput](cfg, models.Cohere_CommandV14)
}
