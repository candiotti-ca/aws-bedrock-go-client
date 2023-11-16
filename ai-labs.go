package awsbedrockgoclient

import (
	"github.com/google/uuid"
)

type AiLabsInput struct {
	Prompt string `json:"prompt"`
	// Use a lower value to decrease randomness in the response
	Temperature float64 `json:"temperature" validate:"min=0,max=1"`
	// Use a lower value to ignore less probable options
	TopP float64 `json:"topP" validate:"min=0,max=1"`
	// Specify the maximum number of tokens to use in the generated response
	MaxTokens int `json:"maxTokens" validate:"min=0,max=8191"`
	// Character sequences that stops generation. Use pipe to separate sequences
	StopSequences []string `json:"stopSequences"`
	// Use a higher value to lower the probability of generating new tokens that already appear at least once in the prompt or in the completion
	CountPenalty struct {
		Scale int `json:"scale" validate:"min=0,max=1"`
	} `json:"countPenalty"`
	// Use a higher value to lower the probability of generating new tokens that already appear at least once in the prompt or in the completion
	PresencePenalty struct {
		Scale float64 `json:"scale" validate:"min=0,max=5"`
	} `json:"presencePenalty"`
	// Use a higher value to lower the probability of generating new tokens that already appear at least once in the prompt or in the completion
	FrequencePenalty struct {
		Scale int `json:"scale" validate:"min=0,max=500"`
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
func NewAiGrandeInstruct(b BedrockRuntime) Client[AiLabsInput, AiLabsOutput] {
	return New[AiLabsInput, AiLabsOutput](b, ai21J2GrandeInstruct)
}

func NewAiJumboInstruct(b BedrockRuntime) Client[AiLabsInput, AiLabsOutput] {
	return New[AiLabsInput, AiLabsOutput](b, ai21J2JumboInstruct)
}

func NewAiJurassic2Mid(b BedrockRuntime) Client[AiLabsInput, AiLabsOutput] {
	return New[AiLabsInput, AiLabsOutput](b, ai21Jurassic2Mid)
}

func NewAiJurassic2MidV1(b BedrockRuntime) Client[AiLabsInput, AiLabsOutput] {
	return New[AiLabsInput, AiLabsOutput](b, ai21Jurassic2Midv1)
}

func NewAiJurassic2Ultra(b BedrockRuntime) Client[AiLabsInput, AiLabsOutput] {
	return New[AiLabsInput, AiLabsOutput](b, ai21Jurassic2Ultra)
}

func NewAiJurassic2UltraV1(b BedrockRuntime) Client[AiLabsInput, AiLabsOutput] {
	return New[AiLabsInput, AiLabsOutput](b, ai21Jurassic2UltraV1)
}
