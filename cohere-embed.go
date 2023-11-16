package awsbedrockgoclient

import (
	"github.com/google/uuid"
)

type CohereEmbedInputType string

const (
	// In search use-cases, use search_document when you encode documents for embeddings that you store in a vector database
	CohereEmbedInputType_SEARCH_DOCUMENT CohereEmbedInputType = "search_document"
	// Use search_query when querying your vector DB to find relevant documents
	CohereEmbedInputType_SEARCH_QUERY CohereEmbedInputType = "search_query"
	// Use classification when using embeddings as an input to a text classifier
	CohereEmbedInputType_CLASSIFICATION CohereEmbedInputType = "classification"
	// Use clustering to cluster the embeddings
	CohereEmbedInputType_CLUSTERING CohereEmbedInputType = "clustering"
)

type CohereEmbedTruncate string

const (
	// (Default) Returns an error when the input exceeds the maximum input token length
	CohereEmbedTruncate_NONE CohereEmbedTruncate = "NONE"
	// Discard the start of the input
	CohereEmbedTruncate_LEFT CohereEmbedTruncate = "LEFT"
	// Discards the end of the input
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

func NewCohereEmbedEnglishV3(b BedrockRuntime) Client[CohereEmbedInput, CohereEmbedOutput] {
	return New[CohereEmbedInput, CohereEmbedOutput](b, cohereEmbedEnglishV3)
}

func NewCohereEmbedMultiV3(b BedrockRuntime) Client[CohereEmbedInput, CohereEmbedOutput] {
	return New[CohereEmbedInput, CohereEmbedOutput](b, cohereEmbedMultiV3)
}
