package awsbedrockgoclient

import (
	"aws-bedrock-go-client/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/google/uuid"
)

type AiLabsInput struct {
	Prompt        string   `json:"prompt"`
	Temperature   float64  `json:"temperature"`
	TopP          float64  `json:"topP"`
	MaxTokens     int      `json:"maxTokens"`
	StopSequences []string `json:"stopSequences"`
	CountPenalty  struct {
		Scale int `json:"scale"`
	} `json:"countPenalty"`
	PresencePenalty struct {
		Scale float64 `json:"scale"`
	} `json:"presencePenalty"`
	FrequencePenalty struct {
		Scale int `json:"scale"`
	} `json:"frequencePenalty"`
}

type AiLabsOutput struct {
	ID     uuid.UUID `json:"id"`
	Prompt struct {
		Text   string `json:"text"`
		Tokens []struct {
			GeneratedToken struct {
				Token      string `json:"token"`
				LogProb    int    `json:"logprob"`
				RawLogProb int    `json:"raw_logprob"`
			} `json:"generatedToken"`
			TopTokens string `json:"topTokens"`
			TextRange struct {
				Start int `json:"start"`
				End   int `json:"end"`
			} `json:"textRange"`
		} `json:"tokens"`
	} `json:"prompt"`
	Completions []struct {
		Data struct {
			Text   string `json:"text"`
			Tokens []struct {
				GeneratedToken struct {
					Token      string `json:"token"`
					LogProb    int    `json:"logprob"`
					RawLogProb int    `json:"raw_logprob"`
				} `json:"generatedToken"`
				TopTokens string `json:"topTokens"`
				TextRange struct {
					Start int `json:"start"`
					End   int `json:"end"`
				} `json:"textRange"`
			} `json:"tokens"`
		} `json:"data"`
		FinishReason struct {
			Reason string `json:"reason"`
			Length int    `json:"length"`
		} `json:"finishReason"`
	} `json:"completions"`
}

// https://docs.ai21.com/reference/j2-complete-ref
func NewAiLabs(cfg aws.Config, model models.Model) Client[AiLabsInput, AiLabsOutput] {
	return New[AiLabsInput, AiLabsOutput](cfg, model)
}