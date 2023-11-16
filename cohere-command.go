package awsbedrockgoclient

import (
	"github.com/google/uuid"
)

type CohereCommandReturnLikelihoods string

const (
	// Only return likelihoods for generated tokens
	CohereCommandReturnLikelihoods_GENERATION CohereCommandReturnLikelihoods = "GENERATION"
	// Return likelihoods for all tokens
	CohereCommandReturnLikelihoods_ALL CohereCommandReturnLikelihoods = "ALL"
	// (Default) Don't return any likelihoods
	CohereCommandReturnLikelihoods_NONE CohereCommandReturnLikelihoods = "NONE"
)

type CohereCommandTruncate string

const (
	// Returns an error when the input exceeds the maximum input token length
	CohereCommandTruncate_NONE CohereCommandTruncate = "NONE"
	// Discard the start of the input
	CohereCommandTruncate_START CohereCommandTruncate = "START"
	// (Default) Discards the end of the input
	CohereCommandTruncate_END CohereCommandTruncate = "END"
)

type CohereCommandInput struct {
	Prompt string `json:"prompt"`
	// (Optional) Use a lower value to decrease randomness in the response
	Temperature float64 `json:"temperature" validate:"min=0,max=5"`
	// (Optional) Use a lower value to ignore less probable options. Set to 0 or 1.0 to disable. If both p and k are enabled, p acts after k
	TopP float64 `json:"p" validate:"min=0,max=1"`
	// (Optional) Specify the number of token choices the model uses to generate the next token. If both p and k are enabled, p acts after k
	TopK             float64 `json:"k" validate:"min=0,max=500"`
	MaxReponseLength int     `json:"max_tokens" validate:"min=1,max=4096"`
	// (Optional) Configure up to four sequences that the model recognizes
	StopSequences []string `json:"stop_sequences"`
	// (optional) Specify how and if the token likelihoods are returned with the response
	ReturnLikelihoods CohereCommandReturnLikelihoods `json:"return_likelihoods"`
	Stream            bool                           `json:"stream"`
	NumGenerations    int                            `json:"num_generations" validate:"min=1,max=5"`
	// (Optional) prevents the model from generating unwanted tokens or incentivizes the model to include desired tokens
	LogitBias map[string]float64 `json:"logit_bias" validate:"min=-10,max=10"`
	// (Optional) Specifies how the API handles inputs longer than the maximum token length
	Truncate CohereCommandTruncate `json:"truncate"`
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

func NewCohereCommandV14(b BedrockRuntime) Client[CohereCommandInput, CohereCommandOutput] {
	return New[CohereCommandInput, CohereCommandOutput](b, cohereCommandV14)
}
