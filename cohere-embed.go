package awsbedrockgoclient

import (
	"aws-bedrock-go-client/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/google/uuid"
)

type CohereEmbedInputType string

const (
	CohereEmbedInputType_SEARCH_DOCUMENT CohereEmbedInputType = "search_document"
	CohereEmbedInputType_SEARCH_QUERY    CohereEmbedInputType = "search_query"
	CohereEmbedInputType_CLASSIFICATION  CohereEmbedInputType = "classification"
	CohereEmbedInputType_CLUSTERING      CohereEmbedInputType = "clustering"
)

type CohereEmbedTruncate string

const (
	CohereEmbedTruncate_NONE  CohereEmbedTruncate = "NONE"
	CohereEmbedTruncate_LEFT  CohereEmbedTruncate = "LEFT"
	CohereEmbedTruncate_RIGHT CohereEmbedTruncate = "RIGHT"
)

type CohereEmbedInput struct {
	Texts     []string            `json:"texts"`
	InputType float64             `json:"input_type"`
	Truncate  CohereEmbedTruncate `json:"truncate"`
}

type CohereEmbedOutput struct {
	ID         uuid.UUID `json:"id"`
	Texts      string    `json:"texts"`
	Embeddings []float64 `json:"embeddings"`
}

func NewCohereEmbed(cfg aws.Config, model models.Model) Client[CohereEmbedInput, CohereEmbedOutput] {
	return New[CohereEmbedInput, CohereEmbedOutput](cfg, model)
}
