package awsbedrockgoclient

type Model string

const (
	cohereCommandV14     Model = "cohere.command-text-v14"
	cohereEmbedEnglishV3 Model = "cohere.embed-english-v3"
	cohereEmbedMultiV3   Model = "cohere.embed-multilingual-v3"

	amazonTitanTextLarge        Model = "amazon.titan-tg1-large"
	amazonTitanTextEmbeddingsV1 Model = "amazon.titan-embed-text-v1"
	amazonTitanTextEmbeddingsV2 Model = "amazon.titan-embed-g1-text-02"
	amazonTitanTextLiteV1       Model = "amazon.titan-text-lite-v1"
	amazonTitanTextExpressV1    Model = "amazon.titan-text-express-v1"

	stabilityAIStableDiffusionXLV0 Model = "stability.stable-diffusion-xl-v0"
	stabilityAIStableDiffusionXLV1 Model = "stability.stable-diffusion-xl-v1"

	ai21J2GrandeInstruct Model = "ai21.j2-grande-instruct"
	ai21J2JumboInstruct  Model = "ai21.j2-jumbo-instruct"
	ai21Jurassic2Mid     Model = "ai21.j2-mid"
	ai21Jurassic2Midv1   Model = "ai21.j2-mid-v1"
	ai21Jurassic2Ultra   Model = "ai21.j2-ultra"
	ai21Jurassic2UltraV1 Model = "ai21.j2-ultra-v1"

	anthropicClaudeInstantV1 Model = "anthropic.claude-instant-v1"
	anthropicClaudeV1        Model = "anthropic.claude-v1"
	anthropicClaudeV2        Model = "anthropic.claude-v2"

	metaLlama2ChatV1 Model = "meta.llama2-13b-chat-v1"
)
